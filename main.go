package main

import (
	"flag"
)

var (
	// Username is your ProtonMail username/email address.
	Username string
	// Password is your ProtonMail password.
	Password string
)

func init() {
	// Initialize flags.
	flag.StringVar(&Username, "username", "", "ProtonMail username")
	flag.StringVar(&Password, "password", "", "ProtonMail password")
}

func main() {
	flag.Parse()
	Count()
}
