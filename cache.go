package main

import (
	"encoding/json"
	"os"
)

var (
	Cache RefreshCache
)

// RefreshCache caches ProtonMail refresh credentials.
type RefreshCache struct {
	// UID ...
	UID string
	// RefreshToken ...
	RefreshToken string
	// Expires stores the UNIX time when these credentials will expire.
	Expires int64
}

// SaveCache persists the cache to disk.
func SaveCache(uid, refreshToken string, expires int64) {
	json, _ := json.Marshal(RefreshCache{UID: uid, RefreshToken: refreshToken, Expires: expires})
	os.WriteFile(".protoncheck", json, os.ModePerm)
}

// LoadCache attempts to load a prior cache into memory.
func LoadCache() error {
	data, err := os.ReadFile(".protoncheck")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Cache)
}
