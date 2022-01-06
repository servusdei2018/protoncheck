package main

import (
	"errors"
	"os"
	"testing"
)

func TestLoadCache(t *testing.T) {
	t.Run("CacheExists", func(t *testing.T) {
		os.Create(".protoncheck")
		got := LoadCache()
		if got != nil {
			t.Errorf("got %s, wanted nil", got)
		}
		os.Remove(".protoncheck")
	})

	t.Run("CacheDoesNotExist", func(t *testing.T) {
		os.Remove(".protoncheck")
		got := LoadCache()
		want := os.ErrNotExist.Error()
		if !errors.Is(got, os.ErrNotExist) {
			t.Errorf("got %s, wanted %s", got, want)
		}
	})
}

func TestSaveCache(t *testing.T) {
	SaveCache("uid", "token", 0)
	_, got := os.Stat(".protoncheck")
	if got != nil {
		t.Errorf("got %s, wanted nil", got)
	}
	os.Remove(".protoncheck")
}
