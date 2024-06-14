// cache.go module contais all logic required to us a cache for any instance.
package api

import "sync"

// -----------------------------------------------------------------------------
//
// ICache
//
// -----------------------------------------------------------------------------

type ICache interface {
	Clear()
	Delete(string)
	Exists(string) bool
	Get(string) (any, bool)
	Set(string, any)
}

// -----------------------------------------------------------------------------
//
// Cache
//
// -----------------------------------------------------------------------------

// Cache struct
type Cache struct {
	data map[string]any
	mu   sync.RWMutex
}

// NewCache creates a new cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

// -----------------------------------------------------------------------------
// Cache public methods
// -----------------------------------------------------------------------------

// Clear removes all entries from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]any)
}

// Delete removes a value from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Exists checks if a key exists in the cache
func (c *Cache) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.data[key]
	return exists
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) GetCache() map[string]any {
	return c.data
}

// Set adds or updates a value in the cache
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
