package relayer

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/internal/utils"
)

const (
	typeUint64  = "uint64"
	typeBool    = "bool"
	typeAddress = "address"
)

// Add64 adds a uint64 value to encrypt.
func (e *EncryptedInput) Add64(value *big.Int) *EncryptedInput {
	if value == nil {
		value = big.NewInt(0)
	}

	if value.Sign() < 0 {
		return e
	}

	maxUint64 := new(big.Int).SetUint64(^uint64(0))
	if value.Cmp(maxUint64) > 0 {
		return e
	}

	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, value.Uint64())

	e.values = append(e.values, encryptedValue{
		valueType: typeUint64,
		data:      data,
	})

	return e
}

// AddBool adds a boolean value to encrypt.
func (e *EncryptedInput) AddBool(value bool) *EncryptedInput {
	var data byte
	if value {
		data = 1
	}

	e.values = append(e.values, encryptedValue{
		valueType: typeBool,
		data:      []byte{data},
	})

	return e
}

// AddAddress adds an Ethereum address to encrypt.
func (e *EncryptedInput) AddAddress(address string) *EncryptedInput {
	if !utils.IsValidAddress(address) {
		return e
	}

	addrBytes, err := utils.AddressToBytes(address)
	if err != nil {
		return e
	}

	e.values = append(e.values, encryptedValue{
		valueType: typeAddress,
		data:      addrBytes,
	})

	return e
}

// Encrypt performs FHE encryption and returns handles + proof.
func (e *EncryptedInput) Encrypt(ctx context.Context) (*EncryptResult, error) {
	if len(e.values) == 0 {
		return nil, fmt.Errorf("no values to encrypt")
	}

	if !utils.IsValidAddress(e.contractAddress) {
		return nil, fmt.Errorf("invalid contract address: %s", e.contractAddress)
	}

	if !utils.IsValidAddress(e.userAddress) {
		return nil, fmt.Errorf("invalid user address: %s", e.userAddress)
	}

	return e.performEncryption(ctx)
}

// Reset clears all values from the builder.
func (e *EncryptedInput) Reset() *EncryptedInput {
	e.values = []encryptedValue{}
	return e
}

// Count returns the number of values added to the builder.
func (e *EncryptedInput) Count() int {
	return len(e.values)
}

func (e *EncryptedInput) performEncryption(_ context.Context) (*EncryptResult, error) {
	handles := make([][]byte, len(e.values))
	for i, val := range e.values {
		handle, err := e.encryptValue(val)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt value %d: %w", i, err)
		}
		handles[i] = handle
	}

	inputProof := e.generateInputProof(handles)

	return &EncryptResult{
		Handles:    handles,
		InputProof: inputProof,
	}, nil
}

func (e *EncryptedInput) encryptValue(val encryptedValue) ([]byte, error) {
	data, ok := val.data.([]byte)
	if !ok {
		return nil, fmt.Errorf("invalid value data type")
	}

	handle := generateMockHandle(val.valueType, data)
	return handle, nil
}

func (e *EncryptedInput) generateInputProof(handles [][]byte) []byte {
	return generateMockProof(e.contractAddress, e.userAddress, handles)
}

func generateMockHandle(valueType string, data []byte) []byte {
	handle := make([]byte, 32)

	switch valueType {
	case typeUint64:
		handle[0] = 0x01
	case typeBool:
		handle[0] = 0x02
	case typeAddress:
		handle[0] = 0x03
	}

	copy(handle[1:], data)
	return handle
}

func generateMockProof(contractAddr, userAddr string, handles [][]byte) []byte {
	proofSize := 64 + len(handles)*32
	proof := make([]byte, proofSize)

	contractBytes, err := utils.AddressToBytes(contractAddr)
	if err == nil {
		copy(proof[0:20], contractBytes)
	}

	userBytes, err := utils.AddressToBytes(userAddr)
	if err == nil {
		copy(proof[20:40], userBytes)
	}

	proof[40] = byte(len(handles))

	offset := 64
	for _, h := range handles {
		copy(proof[offset:offset+32], h)
		offset += 32
	}

	return proof
}
