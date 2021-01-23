package cache

import (
	"log"
	"strings"
	"time"

	"../common"
	"../io/postgres"
	"github.com/allegro/bigcache"
)

var cache *bigcache.BigCache

// Initialize sets up the cache and initializes it
func Initialize() {
	if cache != nil {
		log.Println("cache already initialized")
	}
	log.Println("setting up cache")
	artists := postgres.GetAllArtists()
	config := bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        0,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       len(artists),
		StatsEnabled:       false,
		Verbose:            false,
		HardMaxCacheSize:   0,
		Logger:             bigcache.DefaultLogger(),
	}
	newCache, err := bigcache.NewBigCache(config)
	if err != nil {
		log.Fatal("error starting cache")
	}
	cache = newCache
	fillCache(artists)
}

func fillCache(artists []common.Artist) {
	for i := 0; i < len(artists); i++ {
		cache.Set(artists[i].ID, []byte(strings.Join(artists[i].Recommended[:], ",")))
	}
}

// GetAllEntries returns all entries in the cache
func GetAllEntries() []common.CacheEntry {
	iterator := cache.Iterator()
	iterator.SetNext()
	var entries []common.CacheEntry
	for {
		entryInfo, err := iterator.Value()
		if err != nil {
			panic(err)
		}
		entries = append(entries, common.CacheEntry{
			Key:   entryInfo.Key(),
			Value: string(entryInfo.Value()[:]),
		})
		if !iterator.SetNext() {
			break
		}
	}
	return entries
}
