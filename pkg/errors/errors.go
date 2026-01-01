package errors

import "errors"

var (
	ErrInvalidConfig    = errors.New("invalid configuration")
	ErrInvalidAddress   = errors.New("invalid Ethereum address")
	ErrEncryptionFailed = errors.New("encryption failed")
	ErrNetworkError     = errors.New("network error")
	ErrRelayerError     = errors.New("relayer error")
	ErrNotInitialized   = errors.New("FHEVM instance not initialized")
	ErrInvalidValue     = errors.New("invalid value for encryption")
	ErrTimeout          = errors.New("operation timeout")
)
