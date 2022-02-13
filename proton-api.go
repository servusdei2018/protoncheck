package main

import (
	"context"
	"fmt"

	api "github.com/ProtonMail/proton-bridge/pkg/pmapi"
)

// UserAgentStr represents the ProtonMail user agent.
const UserAgentStr = "ProtonMail/1.13.40 (Android 29; Google Marlin)"

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

// Count retrieves the number of unread messages in a ProtonMail Inbox.
func Count() {
	var client api.Client
	var err error

	// Ensure Username and Password are set.
	if Username == "" || Password == "" {
		fmt.Println("error: Username and Password must be set")
		return
	}

	// Attempt to login with cached refresh credentials.
	if LoadCache() == nil {
		var refresh *api.AuthRefresh
		client, refresh, err = manager.NewClientWithRefresh(context.Background(), Cache.UID, Cache.RefreshToken)
		if err != nil && refresh != nil {
			SaveCache(refresh.UID, refresh.RefreshToken, refresh.ExpiresIn)
		}
	}

	// If the cached credentials failed, login with username and password.
	if err != nil {
		var auth *api.Auth
		client, auth, err = manager.NewClientWithLogin(context.Background(), Username, []byte(Password))
		if err != nil {
			fmt.Println(err)
			return
		}
		SaveCache(auth.UID, auth.RefreshToken, auth.ExpiresIn)
	}

	// Get message counts.
	counts, err := client.CountMessages(context.Background(), "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(counts[0].Unread)
}

// UserAgent returns a User-Agent string to be used in a User-Agent Header.
func UserAgent() string {
	return UserAgentStr
}
