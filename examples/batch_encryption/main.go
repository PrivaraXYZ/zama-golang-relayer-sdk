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
	fmt.Println("  input, err := instance.CreateEncryptedInput(contractAddr, userAddr)")
	fmt.Println("  if err != nil {")
	fmt.Println("      log.Fatal(err)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  // Add multiple values (type-safe, no errors for valid types)")
	fmt.Println("  input.AddUint64(1000)  // Amount")
	fmt.Println("  input.AddUint64(500)   // Price")
	fmt.Println("  input.AddBool(true)    // Is active")
	fmt.Println()
	fmt.Println("  // For address, create value object first")
	fmt.Println("  recipient, err := relayer.NewAddress(recipientAddr)")
	fmt.Println("  if err != nil {")
	fmt.Println("      log.Fatal(err)")
	fmt.Println("  }")
	fmt.Println("  input.AddAddress(recipient)  // Recipient")
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
	fmt.Println("Supported Types (DDD Approach):")
	fmt.Println("  - AddUint64(uint64)        : Type-safe, no validation needed")
	fmt.Println("  - AddBool(bool)            : Type-safe, no validation needed")
	fmt.Println("  - AddAddress(Address)      : Validated value object")
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
