package main

import (
	"encoding/json"
	"os"
)

const (
	// CacheName configures the name of the Cache file.
	CacheName = ".protoncheck"
)

var (
	// Cache ...
	Cache RefreshCache
)

// RefreshCache caches ProtonMail refresh credentials.
type RefreshCache struct {
	// UID ...
	UID string `json:"omitempty,"`
	// RefreshToken ...
	RefreshToken string `json:"omitempty,"`
	// Expires stores the UNIX time when these credentials will expire.
	Expires int64 `json:"omitempty"`
}

// SaveCache persists the cache to disk.
func SaveCache(uid, refreshToken string, expires int64) {
	json, _ := json.Marshal(RefreshCache{UID: uid, RefreshToken: refreshToken, Expires: expires})
	os.WriteFile(CacheName, json, os.ModePerm)
}

// LoadCache attempts to load a prior cache into memory.
func LoadCache() error {
	data, err := os.ReadFile(CacheName)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Cache)
}
