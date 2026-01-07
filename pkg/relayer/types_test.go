package relayer

import (
	"context"
	"testing"
	"time"
)

func TestEncryptedInput_DomainOperations(t *testing.T) {
	contractAddr, _ := NewAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5")
	userAddr, _ := NewAddress("0x1234567890123456789012345678901234567890")

	input := &EncryptedInput{
		contractAddress: contractAddr,
		userAddress:     userAddr,
		values:          []encryptedValue{},
	}

	input.AddUint64(123)
	input.AddBool(true)

	addr, _ := NewAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	input.AddAddress(addr)

	if input.Count() != 3 {
		t.Errorf("Expected 3 values, got %d", input.Count())
	}
}

func TestEncryptedInput_AddUint64(t *testing.T) {
	input := &EncryptedInput{values: []encryptedValue{}}

	input.AddUint64(12345)
	if input.Count() != 1 {
		t.Errorf("Expected 1 value, got %d", input.Count())
	}

	input.AddUint64(0)
	if input.Count() != 2 {
		t.Errorf("Expected 2 values, got %d", input.Count())
	}

	// Type safety: uint64 cannot be negative - this is guaranteed by the type system
	input.AddUint64(18446744073709551615) // max uint64
	if input.Count() != 3 {
		t.Errorf("Expected 3 values, got %d", input.Count())
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

	addr, err := NewAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		t.Fatalf("Failed to create address: %v", err)
	}

	input.AddAddress(addr)
	if input.Count() != 1 {
		t.Errorf("Expected 1 value, got %d", input.Count())
	}
}

func TestAddress_ValueObject(t *testing.T) {
	// Test valid address creation
	addr, err := NewAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		t.Errorf("Valid address should not return error: %v", err)
	}

	if addr.String() != "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" {
		t.Errorf("Address string representation incorrect")
	}

	// Test invalid address creation
	_, err = NewAddress("invalid")
	if err == nil {
		t.Error("Invalid address should return error")
	}

	// Test address equality
	addr2, _ := NewAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if !addr.Equals(addr2) {
		t.Error("Equal addresses should return true")
	}
}

func TestEncryptedInput_Reset(t *testing.T) {
	input := &EncryptedInput{values: []encryptedValue{}}

	input.AddUint64(123)
	input.AddBool(true)
	if input.Count() != 2 {
		t.Errorf("Expected 2 values, got %d", input.Count())
	}

	input.Reset()
	if input.Count() != 0 {
		t.Errorf("Expected 0 values after reset, got %d", input.Count())
	}
}

func TestEncryptedInput_Encrypt(t *testing.T) {
	contractAddr, _ := NewAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5")
	userAddr, _ := NewAddress("0x1234567890123456789012345678901234567890")

	input := &EncryptedInput{
		contractAddress: contractAddr,
		userAddress:     userAddr,
		values:          []encryptedValue{},
	}

	input.AddUint64(12345)
	input.AddBool(true)

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
	contractAddr, _ := NewAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5")
	userAddr, _ := NewAddress("0x1234567890123456789012345678901234567890")

	input := &EncryptedInput{
		contractAddress: contractAddr,
		userAddress:     userAddr,
		values:          []encryptedValue{},
	}

	_, err := input.Encrypt(context.Background())
	if err == nil {
		t.Error("Expected error for empty values")
	}
}

func TestCreateEncryptedInput_Factory(t *testing.T) {
	// Test factory creates valid entity
	instance := &fhevmInstance{}

	input, err := instance.CreateEncryptedInput(
		"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5",
		"0x1234567890123456789012345678901234567890",
	)

	if err != nil {
		t.Fatalf("Factory should create valid entity: %v", err)
	}

	if input == nil {
		t.Error("Factory should return non-nil instance")
	}

	// Test factory validates contract address
	_, err = instance.CreateEncryptedInput(
		"invalid",
		"0x1234567890123456789012345678901234567890",
	)

	if err == nil {
		t.Error("Factory should reject invalid contract address")
	}

	// Test factory validates user address
	_, err = instance.CreateEncryptedInput(
		"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5",
		"invalid",
	)

	if err == nil {
		t.Error("Factory should reject invalid user address")
	}
}

func TestSepoliaConfig(t *testing.T) {
	config, err := SepoliaConfig()
	if err != nil {
		t.Fatalf("SepoliaConfig() should not return error: %v", err)
	}

	if config.ChainID() != 11155111 {
		t.Errorf("Expected Sepolia chain ID 11155111, got %d", config.ChainID())
	}

	if config.GatewayChainID() != 10901 {
		t.Errorf("Expected gateway chain ID 10901, got %d", config.GatewayChainID())
	}

	if config.ACLContract().String() == "" {
		t.Error("ACL contract address should not be empty")
	}

	if config.RelayerURL() == "" {
		t.Error("Relayer URL should not be empty")
	}

	if config.Timeout() <= 0 {
		t.Error("Timeout should be positive")
	}

	if config.MaxRetries() < 0 {
		t.Error("MaxRetries should be non-negative")
	}
}

func TestSepoliaConfigWithOptions(t *testing.T) {
	// Test with custom timeout
	config, err := SepoliaConfig(
		WithTimeout(60*time.Second),
		WithMaxRetries(5),
	)
	if err != nil {
		t.Fatalf("SepoliaConfig with options failed: %v", err)
	}

	if config.Timeout() != 60*time.Second {
		t.Errorf("Expected timeout 60s, got %v", config.Timeout())
	}

	if config.MaxRetries() != 5 {
		t.Errorf("Expected max retries 5, got %d", config.MaxRetries())
	}
}

func TestSepoliaConfigImmutability(t *testing.T) {
	// Config is immutable - all fields are private and can only be accessed via getters
	// Each call should return a new instance
	config1, _ := SepoliaConfig()
	config2, _ := SepoliaConfig()

	// Configs should have independent values
	// Since fields are private, they cannot be modified after creation
	if config1.ChainID() != config2.ChainID() {
		t.Error("Config instances should have same default values")
	}

	// Test that options create independent configs
	config3, _ := SepoliaConfig(WithTimeout(99 * time.Second))
	if config3.Timeout() == config1.Timeout() {
		t.Error("Config with custom option should have different timeout")
	}
}

func TestMainnetConfig(t *testing.T) {
	_, err := MainnetConfig()
	if err == nil {
		t.Error("MainnetConfig should return error (not yet supported)")
	}
}

func TestNewCustomConfig(t *testing.T) {
	config, err := NewCustomConfig(
		999,
		888,
		"https://custom-relayer.com",
		"https://custom-network.com",
	)

	if err != nil {
		t.Fatalf("NewCustomConfig failed: %v", err)
	}

	if config.ChainID() != 999 {
		t.Errorf("Expected chain ID 999, got %d", config.ChainID())
	}

	if config.GatewayChainID() != 888 {
		t.Errorf("Expected gateway chain ID 888, got %d", config.GatewayChainID())
	}

	if config.RelayerURL() != "https://custom-relayer.com" {
		t.Errorf("Unexpected relayer URL: %s", config.RelayerURL())
	}
}

func TestConfigOptions(t *testing.T) {
	tests := []struct {
		name    string
		option  Option
		wantErr bool
	}{
		{
			name:    "valid timeout",
			option:  WithTimeout(45 * time.Second),
			wantErr: false,
		},
		{
			name:    "invalid timeout (zero)",
			option:  WithTimeout(0),
			wantErr: true,
		},
		{
			name:    "invalid timeout (negative)",
			option:  WithTimeout(-1 * time.Second),
			wantErr: true,
		},
		{
			name:    "valid max retries",
			option:  WithMaxRetries(10),
			wantErr: false,
		},
		{
			name:    "invalid max retries (negative)",
			option:  WithMaxRetries(-1),
			wantErr: true,
		},
		{
			name:    "valid custom relayer URL",
			option:  WithCustomRelayerURL("https://test.com"),
			wantErr: false,
		},
		{
			name:    "invalid custom relayer URL (empty)",
			option:  WithCustomRelayerURL(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SepoliaConfig(tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("SepoliaConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
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
				chainID: 11155111,
			},
			wantErr: true,
		},
		{
			name: "missing chain ID",
			config: &FhevmInstanceConfig{
				relayerURL: "https://relayer.testnet.zama.org",
			},
			wantErr: true,
		},
		{
			name: "valid config",
			config: &FhevmInstanceConfig{
				chainID:    11155111,
				relayerURL: "https://relayer.testnet.zama.org",
			},
			wantErr: false,
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
