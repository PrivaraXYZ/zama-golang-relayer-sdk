package utils

import (
	"testing"
)

func TestIsValidAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{"valid lowercase", "0x742d35cc6634c0532925a3b844bc9e7595f0beb5", true},
		{"valid uppercase", "0x742D35CC6634C0532925A3B844BC9E7595F0BEB5", true},
		{"valid mixed", "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5", true},
		{"missing 0x", "742d35cc6634c0532925a3b844bc9e7595f0beb5", false},
		{"too short", "0x742d35cc6634c0532925a3b844bc9e75", false},
		{"too long", "0x742d35cc6634c0532925a3b844bc9e7595f0beb512345", false},
		{"invalid chars", "0x742d35cc6634c0532925a3b844bc9e7595f0beg5", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidAddress(tt.address)
			if got != tt.want {
				t.Errorf("IsValidAddress(%s) = %v, want %v", tt.address, got, tt.want)
			}
		})
	}
}

func TestToChecksumAddress(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			"vitalik address",
			"0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
			"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			false,
		},
		{
			"all lowercase",
			"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
			"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
			false,
		},
		{
			"invalid address",
			"0xinvalid",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToChecksumAddress(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToChecksumAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToChecksumAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChecksumValid(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{"valid checksum", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045", true},
		{"invalid checksum", "0xd8da6bf26964af9d7eed9e03e53415d37aa96045", false},
		{"invalid address", "0xinvalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsChecksumValid(tt.address)
			if got != tt.want {
				t.Errorf("IsChecksumValid(%s) = %v, want %v", tt.address, got, tt.want)
			}
		})
	}
}

func TestAddressToBytes(t *testing.T) {
	address := "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5"
	bytes, err := AddressToBytes(address)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(bytes) != 20 {
		t.Errorf("Expected 20 bytes, got %d", len(bytes))
	}

	back, err := BytesToAddress(bytes)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	normalized, _ := NormalizeAddress(address)
	if back != normalized {
		t.Errorf("Round trip failed: %s != %s", back, normalized)
	}
}

func TestBytesToAddressInvalid(t *testing.T) {
	_, err := BytesToAddress([]byte{1, 2, 3})
	if err == nil {
		t.Error("Expected error for invalid byte length")
	}
}
