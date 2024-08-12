package pokeapi

import (
	"net/http"
	"time"

	"github.com/lululululu5/pokedexcli/internal/pokecache"
)

// Client
type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

// New Client
func NewClient(timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(timeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}