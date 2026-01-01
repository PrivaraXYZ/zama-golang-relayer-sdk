package relayer

// FhevmInstanceConfig holds network configuration.
type FhevmInstanceConfig struct {
	ACLContractAddress                 string
	KMSContractAddress                 string
	InputVerifierContractAddress       string
	VerifyingContractDecryption        string
	VerifyingContractInputVerification string
	ChainID                            uint64
	GatewayChainID                     uint64
	NetworkURL                         string
	RelayerURL                         string
	Auth                               interface{}
}

// SepoliaConfig is configuration for Ethereum Sepolia testnet.
var SepoliaConfig = &FhevmInstanceConfig{
	ACLContractAddress:                 "0xf0Ffdc93b7E186bC2f8CB3dAA75D86d1930A433D",
	KMSContractAddress:                 "0xbE0E383937d564D7FF0BC3b46c51f0bF8d5C311A",
	InputVerifierContractAddress:       "0xBBC1fFCdc7C316aAAd72E807D9b0272BE8F84DA0",
	VerifyingContractDecryption:        "0x5D8BD78e2ea6bbE41f26dFe9fdaEAa349e077478",
	VerifyingContractInputVerification: "0x483b9dE06E4E4C7D35CCf5837A1668487406D955",
	ChainID:                            11155111,
	GatewayChainID:                     10901,
	NetworkURL:                         "https://ethereum-sepolia-rpc.publicnode.com",
	RelayerURL:                         "https://relayer.testnet.zama.org",
}

// MainnetConfig placeholder - not supported in MVP.
var MainnetConfig *FhevmInstanceConfig = nil
