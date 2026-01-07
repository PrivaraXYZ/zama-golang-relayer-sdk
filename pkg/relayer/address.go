package relayer

import (
	"fmt"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/internal/utils"
	sdkerrors "github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/errors"
)

// Address is a value object representing a validated Ethereum address.
// It ensures the address is always in a valid state.
type Address struct {
	value string
}

// NewAddress creates a new Address value object.
// Returns error if the address is invalid.
func NewAddress(addr string) (Address, error) {
	if !utils.IsValidAddress(addr) {
		return Address{}, fmt.Errorf("%w: %s", sdkerrors.ErrInvalidAddress, addr)
	}

	return Address{value: addr}, nil
}

// String returns the string representation of the address.
func (a Address) String() string {
	return a.value
}

// Bytes converts the address to a byte array.
func (a Address) Bytes() ([]byte, error) {
	return utils.AddressToBytes(a.value)
}

// IsZero returns true if the address is the zero value.
func (a Address) IsZero() bool {
	return a.value == ""
}

// Equals checks if two addresses are equal.
func (a Address) Equals(other Address) bool {
	return a.value == other.value
}
