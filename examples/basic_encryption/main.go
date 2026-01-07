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
	sepoliaConfig, _ := relayer.SepoliaConfig()
	fmt.Println("Sepolia Testnet Configuration:")
	fmt.Printf("  Chain ID:          %d\n", sepoliaConfig.ChainID())
	fmt.Printf("  Gateway Chain ID:  %d\n", sepoliaConfig.GatewayChainID())
	fmt.Printf("  Relayer URL:       %s\n", sepoliaConfig.RelayerURL())
	fmt.Printf("  Network URL:       %s\n", sepoliaConfig.NetworkURL())
	fmt.Printf("  Timeout:           %v\n", sepoliaConfig.Timeout())
	fmt.Printf("  Max Retries:       %d\n", sepoliaConfig.MaxRetries())
	fmt.Println()

	fmt.Println("Contract Addresses:")
	fmt.Printf("  ACL:            %s\n", sepoliaConfig.ACLContract().String())
	fmt.Printf("  KMS:            %s\n", sepoliaConfig.KMSContract().String())
	fmt.Printf("  Input Verifier: %s\n", sepoliaConfig.InputVerifierContract().String())
	fmt.Println()

	fmt.Println("Configuration Options:")
	fmt.Println("  // Custom timeout and retries")
	fmt.Println("  config, err := relayer.SepoliaConfig(")
	fmt.Println("      relayer.WithTimeout(60 * time.Second),")
	fmt.Println("      relayer.WithMaxRetries(5),")
	fmt.Println("  )")
	fmt.Println()

	// Example usage (requires running relayer):
	fmt.Println("Example Usage:")
	fmt.Println()
	fmt.Println("  // Get Sepolia configuration")
	fmt.Println("  config, err := relayer.SepoliaConfig()")
	fmt.Println("  if err != nil {")
	fmt.Println("      log.Fatal(err)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  // Initialize FHEVM instance")
	fmt.Println("  instance, err := relayer.CreateInstance(ctx, config)")
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
	fmt.Println("  // Add values to encrypt (type-safe domain operations)")
	fmt.Println("  input.AddUint64(12345)    // No error - uint64 is always valid")
	fmt.Println("  input.AddBool(true)       // No error - bool is always valid")
	fmt.Println()
	fmt.Println("  // For addresses, create Address value object first")
	fmt.Println(`  addr, err := relayer.NewAddress("0xabcd...")`)
	fmt.Println("  if err != nil {")
	fmt.Println("      log.Fatal(err)")
	fmt.Println("  }")
	fmt.Println("  input.AddAddress(addr)    // No error - Address is validated")
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
