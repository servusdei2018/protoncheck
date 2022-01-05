package main

import (
	"flag"
	"fmt"
)

var (
	// Username is your ProtonMail username/email address.
	Username string
	// Password is your ProtonMail password.
	Password string
)

func init() {
	flag.StringVar(&Username, "username", "", "ProtonMail username")
	flag.StringVar(&Password, "password", "", "ProtonMail password")
}

func main() {
	flag.Parse()
	if Username == "" || Password == "" {
		fmt.Println("error: Username and Password must be set")
		return
	}

	count, err := Counts()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
}
