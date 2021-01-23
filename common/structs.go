package common

type Artist struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Popularity  int      `json:"popularity"`
	Recommended []string `json:"recommended"`
}

// CacheEntry is what is stored in the cache. Key is artist ID,
// value is the recommended artists
type CacheEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
