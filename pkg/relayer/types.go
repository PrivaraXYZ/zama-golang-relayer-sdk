package relayer

import (
	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/network"
)

// FhevmInstance represents a connection to the Zama FHEVM protocol.
// Implementations are safe for concurrent use by multiple goroutines.
type FhevmInstance interface {
	CreateEncryptedInput(contractAddress, userAddress string) (*EncryptedInput, error)
	Close() error
}

// EncryptedInput is a domain entity representing encrypted input for FHE operations.
// It is always in a valid state - addresses are validated at creation time.
//
// EncryptedInput is NOT safe for concurrent use by multiple goroutines.
// Each goroutine should create and use its own EncryptedInput instance.
type EncryptedInput struct {
	contractAddress Address
	userAddress     Address
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
	data      []byte
}
