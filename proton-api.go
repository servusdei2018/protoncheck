package main

import (
	"context"

	api "github.com/ProtonMail/proton-bridge/pkg/pmapi"
)

const (
	VERSION = "2.0.1+integrationtests"
)

var (
	manager api.Manager
)

func init() {
	manager = api.New(api.NewConfig("bridge", VERSION))
}

// Retrieves the number of unread messages in a ProtonMail Inbox.
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
