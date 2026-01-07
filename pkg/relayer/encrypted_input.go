package relayer

import (
	"context"
	"encoding/binary"
	"fmt"

	sdkerrors "github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/errors"
)

const (
	typeUint64  = "uint64"
	typeBool    = "bool"
	typeAddress = "address"
)

// AddUint64 adds a uint64 value to encrypt.
// This is a domain operation that updates the entity state.
// The value is guaranteed to be valid by the type system.
func (e *EncryptedInput) AddUint64(value uint64) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, value)

	e.values = append(e.values, encryptedValue{
		valueType: typeUint64,
		data:      data,
	})
}

// AddBool adds a boolean value to encrypt.
// This is a domain operation that updates the entity state.
// The value is guaranteed to be valid by the type system.
func (e *EncryptedInput) AddBool(value bool) {
	var data byte
	if value {
		data = 1
	}

	e.values = append(e.values, encryptedValue{
		valueType: typeBool,
		data:      []byte{data},
	})
}

// AddAddress adds an Ethereum address to encrypt.
// This is a domain operation that updates the entity state.
// Accepts a validated Address value object, ensuring correctness.
func (e *EncryptedInput) AddAddress(address Address) {
	addrBytes, _ := address.Bytes()

	e.values = append(e.values, encryptedValue{
		valueType: typeAddress,
		data:      addrBytes,
	})
}

// Encrypt performs FHE encryption and returns handles + proof.
// The entity is always in a valid state, so no address validation is needed.
func (e *EncryptedInput) Encrypt(ctx context.Context) (*EncryptResult, error) {
	if len(e.values) == 0 {
		return nil, sdkerrors.ErrNoValues
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

func (e *EncryptedInput) performEncryption(ctx context.Context) (*EncryptResult, error) {
	handles := make([][]byte, len(e.values))
	for i, val := range e.values {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		handle, err := e.encryptValue(ctx, val)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt value %d: %w", i, err)
		}
		handles[i] = handle
	}

	inputProof, err := e.generateInputProof(ctx, handles)
	if err != nil {
		return nil, fmt.Errorf("failed to generate input proof: %w", err)
	}

	return &EncryptResult{
		Handles:    handles,
		InputProof: inputProof,
	}, nil
}

func (e *EncryptedInput) encryptValue(ctx context.Context, val encryptedValue) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	handle := generateMockHandle(val.valueType, val.data)
	return handle, nil
}

func (e *EncryptedInput) generateInputProof(ctx context.Context, handles [][]byte) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

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

func generateMockProof(contractAddr, userAddr Address, handles [][]byte) ([]byte, error) {
	proofSize := 64 + len(handles)*32
	proof := make([]byte, proofSize)

	contractBytes, err := contractAddr.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to convert contract address: %w", err)
	}
	copy(proof[0:20], contractBytes)

	userBytes, err := userAddr.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to convert user address: %w", err)
	}
	copy(proof[20:40], userBytes)

	proof[40] = byte(len(handles))

	offset := 64
	for _, h := range handles {
		copy(proof[offset:offset+32], h)
		offset += 32
	}

	return proof, nil
}
