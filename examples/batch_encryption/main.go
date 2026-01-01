// Example: Batch encryption using the Zama FHEVM Relayer SDK.
//
// This example demonstrates encrypting multiple values in a single
// operation using the chainable API.
package main

import (
	"fmt"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/relayer"
)

func main() {
	fmt.Println("Zama FHEVM Relayer SDK - Batch Encryption Example")
	fmt.Println("==================================================")
	fmt.Println()

	fmt.Printf("SDK Version: %s\n", relayer.GetVersion())
	fmt.Println()

	// Demonstrate the batch encryption API
	fmt.Println("Batch Encryption API:")
	fmt.Println()
	fmt.Println("  // The SDK supports encrypting multiple values in a single call.")
	fmt.Println("  // All values are encrypted together and produce a single input proof.")
	fmt.Println()
	fmt.Println("  input := instance.CreateEncryptedInput(contractAddr, userAddr)")
	fmt.Println()
	fmt.Println("  // Chain multiple Add calls")
	fmt.Println("  input.")
	fmt.Println("      Add64(big.NewInt(1000)).      // Amount")
	fmt.Println("      Add64(big.NewInt(500)).       // Price")
	fmt.Println("      AddBool(true).                 // Is active")
	fmt.Println("      AddAddress(recipientAddr)      // Recipient")
	fmt.Println()
	fmt.Println("  // Single encryption call for all values")
	fmt.Println("  result, err := input.Encrypt(ctx)")
	fmt.Println()
	fmt.Println("  // result.Handles[0] = encrypted amount")
	fmt.Println("  // result.Handles[1] = encrypted price")
	fmt.Println("  // result.Handles[2] = encrypted boolean")
	fmt.Println("  // result.Handles[3] = encrypted address")
	fmt.Println("  // result.InputProof = proof for all values")
	fmt.Println()
	fmt.Println()
	fmt.Println("Supported Types:")
	fmt.Println("  - Add64(value)     : Encrypt uint64 (0 to 2^64-1)")
	fmt.Println("  - AddBool(value)   : Encrypt boolean (true/false)")
	fmt.Println("  - AddAddress(addr) : Encrypt Ethereum address")
	fmt.Println()

	fmt.Println("Smart Contract Integration:")
	fmt.Println()
	fmt.Println("  // Pass encrypted values to your FHE-enabled contract")
	fmt.Println("  tx, err := contract.ProcessEncrypted(")
	fmt.Println("      result.Handles[0],  // einput amount")
	fmt.Println("      result.Handles[1],  // einput price")
	fmt.Println("      result.Handles[2],  // einput isActive")
	fmt.Println("      result.Handles[3],  // einput recipient")
	fmt.Println("      result.InputProof,  // Combined proof")
	fmt.Println("  )")
}
