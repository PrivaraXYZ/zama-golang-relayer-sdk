package crypto

import (
	"testing"
	"time"
)

func TestKeyCache_SetGet(t *testing.T) {
	cache := NewKeyCache()

	key := []byte("test-public-key")
	cache.Set(11155111, key)

	got, ok := cache.Get(11155111)
	if !ok {
		t.Error("Expected to find cached key")
	}

	if string(got) != string(key) {
		t.Errorf("Expected %s, got %s", key, got)
	}
}

func TestKeyCache_Expiration(t *testing.T) {
	cache := NewKeyCacheWithOptions(100*time.Millisecond, 100)

	key := []byte("test-key")
	cache.Set(1, key)

	got, ok := cache.Get(1)
	if !ok {
		t.Error("Expected to find cached key immediately")
	}
	if string(got) != string(key) {
		t.Errorf("Expected %s, got %s", key, got)
	}

	time.Sleep(150 * time.Millisecond)

	_, ok = cache.Get(1)
	if ok {
		t.Error("Expected cache entry to be expired")
	}
}

func TestKeyCache_MaxSize(t *testing.T) {
	cache := NewKeyCacheWithOptions(1*time.Hour, 3)

	cache.Set(1, []byte("key1"))
	cache.Set(2, []byte("key2"))
	cache.Set(3, []byte("key3"))

	if cache.Size() != 3 {
		t.Errorf("Expected size 3, got %d", cache.Size())
	}

	cache.Set(4, []byte("key4"))

	if cache.Size() != 3 {
		t.Errorf("Expected size 3 after eviction, got %d", cache.Size())
	}
}

func TestKeyCache_Clear(t *testing.T) {
	cache := NewKeyCache()

	cache.Set(1, []byte("key1"))
	cache.Set(2, []byte("key2"))

	if cache.Size() != 2 {
		t.Errorf("Expected size 2, got %d", cache.Size())
	}

	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", cache.Size())
	}
}

func TestKeyCache_NotFound(t *testing.T) {
	cache := NewKeyCache()

	_, ok := cache.Get(999)
	if ok {
		t.Error("Expected not to find non-existent key")
	}
}
