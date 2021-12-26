package main

import (
	"flag"
	"fmt"
)

var (
	Username string
	Password string
)

func init() {
	flag.StringVar(&Username, "username", "", "ProtonMail username")
	flag.StringVar(&Password, "password", "", "ProtonMail password")
	flag.Parse()
}

func main() {
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
