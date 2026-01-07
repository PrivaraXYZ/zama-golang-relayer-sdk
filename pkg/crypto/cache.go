package crypto

import (
	"sync"
	"time"
)

// KeyCache provides thread-safe caching for public keys.
// KeyCache is safe for concurrent use by multiple goroutines.
type KeyCache struct {
	mu      sync.RWMutex
	keys    map[uint64]cacheEntry
	ttl     time.Duration
	maxSize int
}

type cacheEntry struct {
	key       []byte
	expiresAt time.Time
}

// NewKeyCache creates a new key cache with default settings.
func NewKeyCache() *KeyCache {
	return &KeyCache{
		keys:    make(map[uint64]cacheEntry),
		ttl:     1 * time.Hour,
		maxSize: 100,
	}
}

// NewKeyCacheWithOptions creates a key cache with custom settings.
func NewKeyCacheWithOptions(ttl time.Duration, maxSize int) *KeyCache {
	return &KeyCache{
		keys:    make(map[uint64]cacheEntry),
		ttl:     ttl,
		maxSize: maxSize,
	}
}

// Get retrieves a cached key if it exists and is not expired.
func (c *KeyCache) Get(chainID uint64) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.keys[chainID]
	if !ok {
		return nil, false
	}

	if time.Now().After(entry.expiresAt) {
		return nil, false
	}

	return entry.key, true
}

// Set stores a key in the cache.
func (c *KeyCache) Set(chainID uint64, key []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.keys) >= c.maxSize {
		c.evictOldest()
	}

	c.keys[chainID] = cacheEntry{
		key:       key,
		expiresAt: time.Now().Add(c.ttl),
	}
}

// Clear removes all entries from the cache.
func (c *KeyCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.keys = make(map[uint64]cacheEntry)
}

// Size returns the number of entries in the cache.
func (c *KeyCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.keys)
}

func (c *KeyCache) evictOldest() {
	var oldestChainID uint64
	var oldestTime time.Time

	for chainID, entry := range c.keys {
		if oldestTime.IsZero() || entry.expiresAt.Before(oldestTime) {
			oldestTime = entry.expiresAt
			oldestChainID = chainID
		}
	}

	delete(c.keys, oldestChainID)
}
