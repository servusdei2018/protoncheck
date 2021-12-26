package main

import (
	"context"

	api "github.com/ProtonMail/proton-bridge/pkg/pmapi"
)

var (
	// manager represents a ProtonMail client manager.
	manager api.Manager
)

func init() {
	// Initialize a new ProtonMail client manager.
	config := api.NewConfig("Android", "1.13.39")
	config.GetUserAgent = UserAgent
	config.AppVersion = "Android_1.13.39"
	manager = api.New(config)
}

// Counts retrieves the number of unread messages in a ProtonMail Inbox.
func Counts() (int, error) {
	// Log in to the ProtonMail API.
	client, _, err := manager.NewClientWithLogin(context.Background(), Username, []byte(Password))
	if err != nil {
		return 0, err
	}

	// Get message counts.
	counts, err := client.CountMessages(context.Background(), "")
	if err != nil {
		return 0, err
	}

	return counts[0].Unread, nil
}

// UserAgent returns a User-Agent string to be used in a User-Agent Header.
func UserAgent() string {
	return "ProtonMail/1.13.39 (Android 29; Google Marlin)"
}
