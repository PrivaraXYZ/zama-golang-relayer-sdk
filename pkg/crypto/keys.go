package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/PrivaraXYZ/zama-golang-relayer-sdk/pkg/network"
)

// PublicKeyManager handles fetching and caching of FHEVM public keys.
// PublicKeyManager is safe for concurrent use by multiple goroutines.
type PublicKeyManager struct {
	client *network.Client
	cache  *KeyCache
	mu     sync.RWMutex
}

// NewPublicKeyManager creates a new public key manager.
func NewPublicKeyManager(client *network.Client) *PublicKeyManager {
	return &PublicKeyManager{
		client: client,
		cache:  NewKeyCache(),
	}
}

// GetPublicKey retrieves the public key for a given chain ID.
func (m *PublicKeyManager) GetPublicKey(ctx context.Context, chainID uint64) ([]byte, error) {
	m.mu.RLock()
	cached, ok := m.cache.Get(chainID)
	m.mu.RUnlock()

	if ok {
		return cached, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if cached, ok := m.cache.Get(chainID); ok {
		return cached, nil
	}

	key, err := m.fetchPublicKey(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public key for chain %d: %w", chainID, err)
	}

	m.cache.Set(chainID, key)
	return key, nil
}

func (m *PublicKeyManager) fetchPublicKey(ctx context.Context, chainID uint64) ([]byte, error) {
	path := fmt.Sprintf("%s?chainId=%d", network.EndpointPublicKey, chainID)

	resp, err := m.client.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public key: %w", err)
	}

	var pkResp network.PublicKeyResponse
	if err := json.Unmarshal(resp, &pkResp); err != nil {
		return nil, fmt.Errorf("failed to parse public key response: %w", err)
	}

	return []byte(pkResp.PublicKey), nil
}

// ClearCache clears the public key cache.
func (m *PublicKeyManager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache.Clear()
}

// Close closes the key manager and releases any resources.
// It clears the cache and closes the underlying client.
func (m *PublicKeyManager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Clear()

	if m.client != nil {
		return m.client.Close()
	}

	return nil
}
