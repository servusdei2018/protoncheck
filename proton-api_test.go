package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("NoUsernameOrPassword", func(t *testing.T) {
		Count()
		// Output: error: Username and Password must be set
	})

	t.Run("WithCache", func(t *testing.T) {
		Username = "user"
		Password = "pass"
		SaveCache("uid", "token", 0)
		Count()
		// Output: Incorrect login credentials. Please try again
	})
}

func TestUserAgent(t *testing.T) {
	got := UserAgent()
	want := UserAgentStr
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
