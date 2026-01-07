package errors

import "errors"

// Sentinel errors for the SDK.
// Use errors.Is() to check for these errors in client code.
var (
	// Configuration errors
	ErrInvalidConfig = errors.New("invalid configuration")

	// Address validation errors
	ErrInvalidAddress = errors.New("invalid Ethereum address")

	// Encryption errors
	ErrEncryptionFailed = errors.New("encryption failed")
	ErrInvalidValue     = errors.New("invalid value for encryption")
	ErrNoValues         = errors.New("no values to encrypt")

	// Network errors
	ErrNetworkError = errors.New("network error")
	ErrRelayerError = errors.New("relayer error")
	ErrTimeout      = errors.New("operation timeout")

	// Instance errors
	ErrNotInitialized = errors.New("FHEVM instance not initialized")
)
