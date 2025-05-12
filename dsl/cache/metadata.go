package cache

import (
	"time"
)

type CacheMetadata struct {
	AccessTime time.Time //last time the cache entry was accessed.
	FirstStore time.Time //first time the cache entry was accessed.
	LastMod    time.Time //last time the cache entry was modified.
	TimesUsed  uint32
}

// if 'meta' is eligible to be discarded from cache, return true.
// return false otherwise.
func (meta CacheMetadata) isStale() bool {

	switch _configs.Strategy {
	case "TTL":
		if time.Since(meta.FirstStore) > time.Duration(_configs.TTL) {
			return true
		}

	case "LRU":
		if time.Since(meta.AccessTime) > time.Duration(_configs.TTL) {
			return true
		}

	}

	return false
}
