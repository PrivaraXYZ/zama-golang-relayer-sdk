// Package relayer provides the main SDK for interacting with Zama's FHEVM protocol.
//
// This package implements the core FhevmInstance and EncryptedInput types for
// encrypting values to be used in FHE-enabled smart contracts.
//
// Example usage:
//
//	import "github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/relayer"
//
//	instance, err := relayer.CreateInstance(ctx, relayer.SepoliaConfig)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	input := instance.CreateEncryptedInput(contractAddr, userAddr)
//	input.Add64(big.NewInt(12345))
//	result, err := input.Encrypt(ctx)
package relayer
