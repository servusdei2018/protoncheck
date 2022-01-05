package main

import "testing"

func TestUserAgent(t *testing.T) {
	got := UserAgent()
	want := UserAgentStr
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
