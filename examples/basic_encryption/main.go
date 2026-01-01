// Example: Basic encryption using the Zama FHEVM Relayer SDK.
//
// This example demonstrates the SDK API and configuration.
// Note: Actual encryption requires a running Zama relayer service.
package main

import (
	"fmt"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/relayer"
)

func main() {
	fmt.Println("Zama FHEVM Relayer SDK - Basic Encryption Example")
	fmt.Println("==================================================")
	fmt.Println()

	// Display SDK version
	fmt.Printf("SDK Version: %s\n", relayer.GetVersion())
	fmt.Println()

	// Show the Sepolia configuration
	fmt.Println("Sepolia Testnet Configuration:")
	fmt.Printf("  Chain ID:          %d\n", relayer.SepoliaConfig.ChainID)
	fmt.Printf("  Gateway Chain ID:  %d\n", relayer.SepoliaConfig.GatewayChainID)
	fmt.Printf("  Relayer URL:       %s\n", relayer.SepoliaConfig.RelayerURL)
	fmt.Printf("  Network URL:       %s\n", relayer.SepoliaConfig.NetworkURL)
	fmt.Println()

	fmt.Println("Contract Addresses:")
	fmt.Printf("  ACL:            %s\n", relayer.SepoliaConfig.ACLContractAddress)
	fmt.Printf("  KMS:            %s\n", relayer.SepoliaConfig.KMSContractAddress)
	fmt.Printf("  Input Verifier: %s\n", relayer.SepoliaConfig.InputVerifierContractAddress)
	fmt.Println()

	// Example usage (requires running relayer):
	fmt.Println("Example Usage:")
	fmt.Println()
	fmt.Println("  // Initialize FHEVM instance")
	fmt.Println("  instance, err := relayer.CreateInstance(ctx, relayer.SepoliaConfig)")
	fmt.Println("  if err != nil {")
	fmt.Println("      log.Fatal(err)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  // Create encrypted input for contract interaction")
	fmt.Println("  input := instance.CreateEncryptedInput(")
	fmt.Println(`      "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb5", // Contract`)
	fmt.Println(`      "0x1234567890123456789012345678901234567890", // User`)
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("  // Add values to encrypt (chainable API)")
	fmt.Println("  input.")
	fmt.Println("      Add64(big.NewInt(12345)).  // Encrypt uint64")
	fmt.Println("      AddBool(true).              // Encrypt boolean")
	fmt.Println(`      AddAddress("0xabcd...")     // Encrypt address`)
	fmt.Println()
	fmt.Println("  // Perform encryption")
	fmt.Println("  result, err := input.Encrypt(ctx)")
	fmt.Println("  if err != nil {")
	fmt.Println("      log.Fatal(err)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  // Use handles and proof in smart contract call")
	fmt.Println("  // Handles slice contains encrypted value handles")
	fmt.Println("  // InputProof contains the cryptographic proof")
	fmt.Println()
	fmt.Println()
	fmt.Println("For full examples, see the SDK documentation.")
}
