package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

// IsValidAddress checks if a string is a valid Ethereum address.
func IsValidAddress(address string) bool {
	if !strings.HasPrefix(address, "0x") {
		return false
	}

	address = strings.TrimPrefix(address, "0x")
	if len(address) != 40 {
		return false
	}

	_, err := hex.DecodeString(address)
	return err == nil
}

// ToChecksumAddress converts an Ethereum address to its checksummed format.
func ToChecksumAddress(address string) (string, error) {
	if !IsValidAddress(address) {
		return "", fmt.Errorf("invalid address: %s", address)
	}

	address = strings.ToLower(strings.TrimPrefix(address, "0x"))

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(address))
	hashBytes := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)

	result := make([]byte, len(address))
	for i, c := range []byte(address) {
		if c >= 'a' && c <= 'f' {
			if hashHex[i] >= '8' {
				result[i] = c - 32
			} else {
				result[i] = c
			}
		} else {
			result[i] = c
		}
	}

	return "0x" + string(result), nil
}

// IsChecksumValid verifies if an address has a valid checksum.
func IsChecksumValid(address string) bool {
	if !IsValidAddress(address) {
		return false
	}

	checksummed, err := ToChecksumAddress(address)
	if err != nil {
		return false
	}

	return address == checksummed
}

// NormalizeAddress normalizes an address to lowercase with 0x prefix.
func NormalizeAddress(address string) (string, error) {
	if !IsValidAddress(address) {
		return "", fmt.Errorf("invalid address: %s", address)
	}

	return strings.ToLower(address), nil
}

// AddressToBytes converts an Ethereum address to a 20-byte array.
func AddressToBytes(address string) ([]byte, error) {
	if !IsValidAddress(address) {
		return nil, fmt.Errorf("invalid address: %s", address)
	}

	address = strings.TrimPrefix(address, "0x")
	return hex.DecodeString(address)
}

// BytesToAddress converts a 20-byte array to an Ethereum address.
func BytesToAddress(b []byte) (string, error) {
	if len(b) != 20 {
		return "", fmt.Errorf("invalid byte length: expected 20, got %d", len(b))
	}

	return "0x" + hex.EncodeToString(b), nil
}
