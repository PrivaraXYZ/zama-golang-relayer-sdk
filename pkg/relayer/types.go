package relayer

import (
	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/network"
)

// FhevmInstance represents a connection to the Zama FHEVM protocol.
type FhevmInstance interface {
	CreateEncryptedInput(contractAddress, userAddress string) *EncryptedInput
}

// EncryptedInput is a builder for collecting values to encrypt.
type EncryptedInput struct {
	contractAddress string
	userAddress     string
	values          []encryptedValue
	config          *FhevmInstanceConfig
	client          *network.Client
	publicKey       []byte
}

// EncryptResult contains encryption output.
type EncryptResult struct {
	Handles    [][]byte
	InputProof []byte
}

type encryptedValue struct {
	valueType string
	data      interface{}
}
