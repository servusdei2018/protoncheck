package main

import (
	"errors"
	"testing"
	"os"
)

func TestLoadCache(t *testing.T) {
	got := LoadCache()
	if got != nil && !errors.Is(err, os.ErrNotExist) {
		t.Errorf("got %s, wanted nil || os.ErrNotExist", got)
	}
}
