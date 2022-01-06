package main

import (
	"errors"
	"os"
	"testing"
)

func TestLoadCache(t *testing.T) {
	t.Run("CacheExists", func(t *testing.T) {
		SaveCache("uid", "token", 0)
		got := LoadCache()
		if got != nil {
			t.Errorf("got %s, wanted nil", got)
		}
		os.Remove(CacheName)
	})

	t.Run("CacheDoesNotExist", func(t *testing.T) {
		os.Remove(CacheName)
		got := LoadCache()
		want := os.ErrNotExist.Error()
		if !errors.Is(got, os.ErrNotExist) {
			t.Errorf("got %s, wanted %s", got, want)
		}
	})
}

func TestSaveCache(t *testing.T) {
	SaveCache("uid", "token", 0)
	_, got := os.Stat(CacheName)
	if got != nil {
		t.Errorf("got %s, wanted nil", got)
	}
	os.Remove(CacheName)
}
