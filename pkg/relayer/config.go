package relayer

import (
	"fmt"
	"time"
)

// FhevmInstanceConfig holds network configuration for FHEVM instance.
// It is immutable after creation and always in a valid state.
// All fields are private to enforce immutability. Use getter methods to access values.
type FhevmInstanceConfig struct {
	// Network identifiers
	chainID        uint64
	gatewayChainID uint64

	// URLs
	networkURL string
	relayerURL string

	// Contract addresses (validated value objects)
	aclContract                        Address
	kmsContract                        Address
	inputVerifierContract              Address
	verifyingContractDecryption        Address
	verifyingContractInputVerification Address

	// HTTP client settings
	timeout    time.Duration
	maxRetries int

	// Optional authentication
	auth interface{}
}

// Getter methods for FhevmInstanceConfig

// ChainID returns the blockchain chain ID.
func (c *FhevmInstanceConfig) ChainID() uint64 {
	return c.chainID
}

// GatewayChainID returns the gateway chain ID.
func (c *FhevmInstanceConfig) GatewayChainID() uint64 {
	return c.gatewayChainID
}

// NetworkURL returns the network RPC URL.
func (c *FhevmInstanceConfig) NetworkURL() string {
	return c.networkURL
}

// RelayerURL returns the relayer URL.
func (c *FhevmInstanceConfig) RelayerURL() string {
	return c.relayerURL
}

// ACLContract returns the ACL contract address.
func (c *FhevmInstanceConfig) ACLContract() Address {
	return c.aclContract
}

// KMSContract returns the KMS contract address.
func (c *FhevmInstanceConfig) KMSContract() Address {
	return c.kmsContract
}

// InputVerifierContract returns the input verifier contract address.
func (c *FhevmInstanceConfig) InputVerifierContract() Address {
	return c.inputVerifierContract
}

// VerifyingContractDecryption returns the decryption verifying contract address.
func (c *FhevmInstanceConfig) VerifyingContractDecryption() Address {
	return c.verifyingContractDecryption
}

// VerifyingContractInputVerification returns the input verification verifying contract address.
func (c *FhevmInstanceConfig) VerifyingContractInputVerification() Address {
	return c.verifyingContractInputVerification
}

// Timeout returns the HTTP request timeout.
func (c *FhevmInstanceConfig) Timeout() time.Duration {
	return c.timeout
}

// MaxRetries returns the maximum number of retry attempts.
func (c *FhevmInstanceConfig) MaxRetries() int {
	return c.maxRetries
}

// Auth returns the authentication credentials.
func (c *FhevmInstanceConfig) Auth() interface{} {
	return c.auth
}

// Option is a functional option for configuring FhevmInstanceConfig.
type Option func(*FhevmInstanceConfig) error

// WithTimeout sets the HTTP request timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *FhevmInstanceConfig) error {
		if timeout <= 0 {
			return fmt.Errorf("timeout must be positive, got: %v", timeout)
		}
		c.timeout = timeout
		return nil
	}
}

// WithMaxRetries sets the maximum number of retry attempts.
func WithMaxRetries(retries int) Option {
	return func(c *FhevmInstanceConfig) error {
		if retries < 0 {
			return fmt.Errorf("max retries cannot be negative, got: %d", retries)
		}
		c.maxRetries = retries
		return nil
	}
}

// WithCustomRelayerURL overrides the default relayer URL.
func WithCustomRelayerURL(url string) Option {
	return func(c *FhevmInstanceConfig) error {
		if url == "" {
			return fmt.Errorf("relayer URL cannot be empty")
		}
		c.relayerURL = url
		return nil
	}
}

// WithCustomNetworkURL overrides the default network RPC URL.
func WithCustomNetworkURL(url string) Option {
	return func(c *FhevmInstanceConfig) error {
		if url == "" {
			return fmt.Errorf("network URL cannot be empty")
		}
		c.networkURL = url
		return nil
	}
}

// WithAuth sets authentication credentials for the relayer.
func WithAuth(auth interface{}) Option {
	return func(c *FhevmInstanceConfig) error {
		c.auth = auth
		return nil
	}
}

// WithChainID sets the blockchain chain ID.
func WithChainID(chainID uint64) Option {
	return func(c *FhevmInstanceConfig) error {
		if chainID == 0 {
			return fmt.Errorf("chain ID cannot be zero")
		}
		c.chainID = chainID
		return nil
	}
}

// WithGatewayChainID sets the gateway chain ID.
func WithGatewayChainID(gatewayChainID uint64) Option {
	return func(c *FhevmInstanceConfig) error {
		if gatewayChainID == 0 {
			return fmt.Errorf("gateway chain ID cannot be zero")
		}
		c.gatewayChainID = gatewayChainID
		return nil
	}
}

// WithACLContract sets the ACL contract address.
func WithACLContract(address string) Option {
	return func(c *FhevmInstanceConfig) error {
		addr, err := NewAddress(address)
		if err != nil {
			return fmt.Errorf("invalid ACL contract address: %w", err)
		}
		c.aclContract = addr
		return nil
	}
}

// WithKMSContract sets the KMS contract address.
func WithKMSContract(address string) Option {
	return func(c *FhevmInstanceConfig) error {
		addr, err := NewAddress(address)
		if err != nil {
			return fmt.Errorf("invalid KMS contract address: %w", err)
		}
		c.kmsContract = addr
		return nil
	}
}

// WithInputVerifierContract sets the input verifier contract address.
func WithInputVerifierContract(address string) Option {
	return func(c *FhevmInstanceConfig) error {
		addr, err := NewAddress(address)
		if err != nil {
			return fmt.Errorf("invalid input verifier contract address: %w", err)
		}
		c.inputVerifierContract = addr
		return nil
	}
}

// WithVerifyingContractDecryption sets the decryption verifying contract address.
func WithVerifyingContractDecryption(address string) Option {
	return func(c *FhevmInstanceConfig) error {
		addr, err := NewAddress(address)
		if err != nil {
			return fmt.Errorf("invalid decryption verifying contract address: %w", err)
		}
		c.verifyingContractDecryption = addr
		return nil
	}
}

// WithVerifyingContractInputVerification sets the input verification verifying contract address.
func WithVerifyingContractInputVerification(address string) Option {
	return func(c *FhevmInstanceConfig) error {
		addr, err := NewAddress(address)
		if err != nil {
			return fmt.Errorf("invalid input verification verifying contract address: %w", err)
		}
		c.verifyingContractInputVerification = addr
		return nil
	}
}

// SepoliaConfig returns a pre-configured setup for Ethereum Sepolia testnet.
// Returns a new instance each time to ensure immutability.
func SepoliaConfig(opts ...Option) (*FhevmInstanceConfig, error) {
	// Create validated Address value objects
	aclAddr, _ := NewAddress("0xf0Ffdc93b7E186bC2f8CB3dAA75D86d1930A433D")
	kmsAddr, _ := NewAddress("0xbE0E383937d564D7FF0BC3b46c51f0bF8d5C311A")
	verifierAddr, _ := NewAddress("0xBBC1fFCdc7C316aAAd72E807D9b0272BE8F84DA0")
	decryptionAddr, _ := NewAddress("0x5D8BD78e2ea6bbE41f26dFe9fdaEAa349e077478")
	inputVerificationAddr, _ := NewAddress("0x483b9dE06E4E4C7D35CCf5837A1668487406D955")

	cfg := &FhevmInstanceConfig{
		chainID:                            11155111,
		gatewayChainID:                     10901,
		networkURL:                         "https://ethereum-sepolia-rpc.publicnode.com",
		relayerURL:                         "https://relayer.testnet.zama.org",
		aclContract:                        aclAddr,
		kmsContract:                        kmsAddr,
		inputVerifierContract:              verifierAddr,
		verifyingContractDecryption:        decryptionAddr,
		verifyingContractInputVerification: inputVerificationAddr,
		timeout:                            30 * time.Second,
		maxRetries:                         3,
	}

	// Apply custom options
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, fmt.Errorf("failed to apply config option: %w", err)
		}
	}

	return cfg, nil
}

// MainnetConfig returns a pre-configured setup for Ethereum mainnet.
// Currently not supported.
func MainnetConfig(opts ...Option) (*FhevmInstanceConfig, error) {
	return nil, fmt.Errorf("mainnet configuration is not yet supported (planned for v0.4.0)")
}

// NewCustomConfig creates a fully custom configuration.
// Use this for custom networks or testnets.
func NewCustomConfig(
	chainID uint64,
	gatewayChainID uint64,
	relayerURL string,
	networkURL string,
	opts ...Option,
) (*FhevmInstanceConfig, error) {
	cfg := &FhevmInstanceConfig{
		chainID:        chainID,
		gatewayChainID: gatewayChainID,
		relayerURL:     relayerURL,
		networkURL:     networkURL,
		timeout:        30 * time.Second,
		maxRetries:     3,
	}

	// Apply custom options
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, fmt.Errorf("failed to apply config option: %w", err)
		}
	}

	// Validate the configuration
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}
