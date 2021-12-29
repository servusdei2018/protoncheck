package main

import (
	"context"
	"fmt"
	"time"

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
	var client api.Client
	var err error

	err = LoadCache()

	// Attempt to login with cached refresh credentials.
	if err == nil {
		var refresh *api.AuthRefresh
		if time.Unix(Cache.Expires, 0).Sub(time.Now()) <= 5*time.Minute {
			err = fmt.Errorf("error: cache expired.")
		}
		client, refresh, err = manager.NewClientWithRefresh(context.Background(), Cache.UID, Cache.RefreshToken)
		SaveCache(refresh.UID, refresh.RefreshToken, refresh.ExpiresIn)
	}

	// Login with username and password.
	if err != nil {
		var auth *api.Auth
		client, auth, err = manager.NewClientWithLogin(context.Background(), Username, []byte(Password))
		if err != nil {
			return 0, err
		}
		SaveCache(auth.UID, auth.RefreshToken, auth.ExpiresIn)
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
