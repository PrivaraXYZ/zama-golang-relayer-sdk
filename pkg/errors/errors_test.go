package errors

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{"InvalidConfig", ErrInvalidConfig},
		{"InvalidAddress", ErrInvalidAddress},
		{"EncryptionFailed", ErrEncryptionFailed},
		{"NetworkError", ErrNetworkError},
		{"InvalidValue", ErrInvalidValue},
		{"NoValues", ErrNoValues},
		{"Timeout", ErrTimeout},
		{"RelayerError", ErrRelayerError},
		{"NotInitialized", ErrNotInitialized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Error("Error should not be nil")
			}

			if errors.Is(tt.err, tt.err) == false {
				t.Error("errors.Is should match")
			}
		})
	}
}
