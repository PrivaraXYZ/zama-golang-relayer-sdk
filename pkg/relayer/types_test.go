package relayer

import (
	"context"
	"math/big"
	"testing"
)

func TestEncryptedInput_Builder(t *testing.T) {
	input := &EncryptedInput{
		contractAddress: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5",
		userAddress:     "0x1234567890123456789012345678901234567890",
		values:          []encryptedValue{},
	}

	input.Add64(big.NewInt(123)).AddBool(true).AddAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")

	if input.Count() != 3 {
		t.Errorf("Expected 3 values, got %d", input.Count())
	}
}

func TestEncryptedInput_Add64(t *testing.T) {
	input := &EncryptedInput{values: []encryptedValue{}}

	input.Add64(big.NewInt(12345))
	if input.Count() != 1 {
		t.Errorf("Expected 1 value, got %d", input.Count())
	}

	input.Add64(big.NewInt(0))
	if input.Count() != 2 {
		t.Errorf("Expected 2 values, got %d", input.Count())
	}

	input.Add64(big.NewInt(-1))
	if input.Count() != 2 {
		t.Errorf("Expected 2 values (negative ignored), got %d", input.Count())
	}
}

func TestEncryptedInput_AddBool(t *testing.T) {
	input := &EncryptedInput{values: []encryptedValue{}}

	input.AddBool(true)
	if input.Count() != 1 {
		t.Errorf("Expected 1 value, got %d", input.Count())
	}

	input.AddBool(false)
	if input.Count() != 2 {
		t.Errorf("Expected 2 values, got %d", input.Count())
	}
}

func TestEncryptedInput_AddAddress(t *testing.T) {
	input := &EncryptedInput{values: []encryptedValue{}}

	input.AddAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if input.Count() != 1 {
		t.Errorf("Expected 1 value, got %d", input.Count())
	}

	input.AddAddress("invalid")
	if input.Count() != 1 {
		t.Errorf("Expected 1 value (invalid ignored), got %d", input.Count())
	}
}

func TestEncryptedInput_Reset(t *testing.T) {
	input := &EncryptedInput{values: []encryptedValue{}}

	input.Add64(big.NewInt(123)).AddBool(true)
	if input.Count() != 2 {
		t.Errorf("Expected 2 values, got %d", input.Count())
	}

	input.Reset()
	if input.Count() != 0 {
		t.Errorf("Expected 0 values after reset, got %d", input.Count())
	}
}

func TestEncryptedInput_Encrypt(t *testing.T) {
	input := &EncryptedInput{
		contractAddress: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5",
		userAddress:     "0x1234567890123456789012345678901234567890",
		values:          []encryptedValue{},
	}

	input.Add64(big.NewInt(12345)).AddBool(true)

	result, err := input.Encrypt(context.Background())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.Handles) != 2 {
		t.Errorf("Expected 2 handles, got %d", len(result.Handles))
	}

	for i, h := range result.Handles {
		if len(h) != 32 {
			t.Errorf("Handle %d should be 32 bytes, got %d", i, len(h))
		}
	}

	if len(result.InputProof) == 0 {
		t.Error("InputProof should not be empty")
	}
}

func TestEncryptedInput_EncryptEmpty(t *testing.T) {
	input := &EncryptedInput{
		contractAddress: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5",
		userAddress:     "0x1234567890123456789012345678901234567890",
		values:          []encryptedValue{},
	}

	_, err := input.Encrypt(context.Background())
	if err == nil {
		t.Error("Expected error for empty values")
	}
}

func TestEncryptedInput_EncryptInvalidAddress(t *testing.T) {
	input := &EncryptedInput{
		contractAddress: "invalid",
		userAddress:     "0x1234567890123456789012345678901234567890",
		values:          []encryptedValue{},
	}

	input.Add64(big.NewInt(123))

	_, err := input.Encrypt(context.Background())
	if err == nil {
		t.Error("Expected error for invalid contract address")
	}
}

func TestSepoliaConfig(t *testing.T) {
	if SepoliaConfig.ChainID != 11155111 {
		t.Errorf("Expected Sepolia chain ID 11155111, got %d", SepoliaConfig.ChainID)
	}

	if SepoliaConfig.ACLContractAddress == "" {
		t.Error("ACL contract address should not be empty")
	}

	if SepoliaConfig.RelayerURL == "" {
		t.Error("Relayer URL should not be empty")
	}
}

func TestVersion(t *testing.T) {
	v := GetVersion()
	if v != "0.1.0-rc.1" {
		t.Errorf("Expected version 0.1.0-rc.1, got %s", v)
	}
}

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  *FhevmInstanceConfig
		wantErr bool
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "missing relayer URL",
			config: &FhevmInstanceConfig{
				ChainID: 11155111,
			},
			wantErr: true,
		},
		{
			name: "missing chain ID",
			config: &FhevmInstanceConfig{
				RelayerURL: "https://relayer.testnet.zama.org",
			},
			wantErr: true,
		},
		{
			name: "valid config",
			config: &FhevmInstanceConfig{
				ChainID:    11155111,
				RelayerURL: "https://relayer.testnet.zama.org",
			},
			wantErr: false,
		},
		{
			name: "invalid ACL address",
			config: &FhevmInstanceConfig{
				ChainID:            11155111,
				RelayerURL:         "https://relayer.testnet.zama.org",
				ACLContractAddress: "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
