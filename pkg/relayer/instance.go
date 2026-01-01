package relayer

import (
	"context"
	"fmt"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/internal/utils"
	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/crypto"
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

	client := network.NewClient(config.RelayerURL)
	keyMgr := crypto.NewPublicKeyManager(client)

	publicKey, err := keyMgr.GetPublicKey(ctx, config.ChainID)
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

// CreateEncryptedInput creates a new EncryptedInput builder.
func (f *fhevmInstance) CreateEncryptedInput(contractAddress, userAddress string) *EncryptedInput {
	return &EncryptedInput{
		contractAddress: contractAddress,
		userAddress:     userAddress,
		values:          []encryptedValue{},
		config:          f.config,
		client:          f.client,
		publicKey:       f.publicKey,
	}
}

func validateConfig(config *FhevmInstanceConfig) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	if config.RelayerURL == "" {
		return fmt.Errorf("relayer URL is required")
	}

	if config.ChainID == 0 {
		return fmt.Errorf("chain ID is required")
	}

	if config.ACLContractAddress != "" && !utils.IsValidAddress(config.ACLContractAddress) {
		return fmt.Errorf("invalid ACL contract address")
	}

	if config.KMSContractAddress != "" && !utils.IsValidAddress(config.KMSContractAddress) {
		return fmt.Errorf("invalid KMS contract address")
	}

	if config.InputVerifierContractAddress != "" && !utils.IsValidAddress(config.InputVerifierContractAddress) {
		return fmt.Errorf("invalid input verifier contract address")
	}

	return nil
}
