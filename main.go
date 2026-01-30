/*
Filename: main.go
Description: routes api + requests badges
Created by: main
        at: 12:12 on Thursday, the 29th of January, 2026.
Last edited 23:32 on Friday, the 30th of January, 2026
*/

package main

import (
	"sync"
	"time"
)

const ttl = 5 * time.Minute

var svgCache cache

func init() {
	svgCache.entries = make(map[string]cachedSVG)

	// run the cleanup every ten minutes
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			expireCached()
		}
	}()

}

func main() {
}

// a generated SVG response has a TTL per project requested, which we don't have to rebuild
type cachedSVG struct {
	SVG     string
	expires time.Time
}

// the cache stores our SVG responses in memory
type cache struct {
	mu      sync.RWMutex
	entries map[string]cachedSVG
}

// load a SVG response if it exists in our cache
func loadCachedSVG(key string) (string, bool) {
	svgCache.mu.RLock()
	entry, exists := svgCache.entries[key]
	svgCache.mu.RUnlock()

	if !exists || time.Now().After(entry.expires) {
		return "", false
	}
	return entry.SVG, true
}

// saves a generated SVG into the cache
func storeCachedSVG(key, svg string) {
	svgCache.mu.Lock()
	svgCache.entries[key] = cachedSVG{
		SVG:     svg,
		expires: time.Now().Add(ttl),
	}
	svgCache.mu.Unlock()
}

// goes through the caches, removing any expired SVGs
func expireCached() {
	svgCache.mu.Lock()
	defer svgCache.mu.Unlock()

	now := time.Now()
	for key, entry := range svgCache.entries {
		if now.After(entry.expires) {
			delete(svgCache.entries, key)
		}
	}
}
