package relayer

import (
	"context"
	"fmt"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/crypto"
	sdkerrors "github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/errors"
	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/network"
)

// fhevmInstance is the concrete implementation of FhevmInstance.
type fhevmInstance struct {
	config    *FhevmInstanceConfig
	client    *network.Client
	keyMgr    *crypto.PublicKeyManager
	publicKey []byte
}

// CreateInstance initializes an FhevmInstance with configuration.
func CreateInstance(ctx context.Context, config *FhevmInstanceConfig) (FhevmInstance, error) {
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	client := network.NewClient(config.RelayerURL())
	keyMgr := crypto.NewPublicKeyManager(client)

	publicKey, err := keyMgr.GetPublicKey(ctx, config.ChainID())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public key: %w", err)
	}

	return &fhevmInstance{
		config:    config,
		client:    client,
		keyMgr:    keyMgr,
		publicKey: publicKey,
	}, nil
}

// CreateEncryptedInput creates a new EncryptedInput entity.
// This is a factory method that validates addresses and ensures the entity is in a valid state.
func (f *fhevmInstance) CreateEncryptedInput(contractAddress, userAddress string) (*EncryptedInput, error) {
	contractAddr, err := NewAddress(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid contract address: %w", err)
	}

	userAddr, err := NewAddress(userAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid user address: %w", err)
	}

	return &EncryptedInput{
		contractAddress: contractAddr,
		userAddress:     userAddr,
		values:          []encryptedValue{},
		config:          f.config,
		client:          f.client,
		publicKey:       f.publicKey,
	}, nil
}

// Close closes the instance and releases all resources.
// It closes the key manager which in turn closes the HTTP client.
// After calling Close, the instance should not be used.
func (f *fhevmInstance) Close() error {
	if f.keyMgr != nil {
		return f.keyMgr.Close()
	}
	return nil
}

func validateConfig(config *FhevmInstanceConfig) error {
	if config == nil {
		return fmt.Errorf("%w: config cannot be nil", sdkerrors.ErrInvalidConfig)
	}

	if config.RelayerURL() == "" {
		return fmt.Errorf("%w: relayer URL is required", sdkerrors.ErrInvalidConfig)
	}

	if config.ChainID() == 0 {
		return fmt.Errorf("%w: chain ID is required", sdkerrors.ErrInvalidConfig)
	}

	return nil
}
