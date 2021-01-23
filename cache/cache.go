package cache

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

var cache *bigcache.BigCache

// Initialize sets up the cache and initializes it
func Initialize() {
	if cache != nil {
		log.Println("cache already initialized")
	}
	log.Println("setting up cache")
	newCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	cache = newCache
	if err != nil {
		log.Fatal("error starting cache")
	}
}
