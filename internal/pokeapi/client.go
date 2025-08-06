package pokeapi

import (
	"github.com/LucasSim0n/Pokedex_in_Go/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
