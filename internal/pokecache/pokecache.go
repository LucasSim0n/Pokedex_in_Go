package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt       time.Time
	val             []byte
	capturedPokemon bool
}

type Cache struct {
	cache map[string]*cacheEntry
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]*cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte, capturedPokemon bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = &cacheEntry{
		createdAt:       time.Now(),
		val:             val,
		capturedPokemon: capturedPokemon,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val, ok := c.cache[key]
	if !ok {
		var zero []byte
		return zero, false
	}
	return val.val, ok
}

func (c *Cache) UpdateCP(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val := c.cache[key]
	val.capturedPokemon = true
}

func (c *Cache) GetCP(key string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	if !ok {
		return false
	}
	return val.capturedPokemon
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) && !v.capturedPokemon {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) GetAllPokemons() [][]byte {
	var pokemons [][]byte

	for _, val := range c.cache {
		if val.capturedPokemon {
			pokemons = append(pokemons, val.val)
		}
	}
	return pokemons
}
