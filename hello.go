package main

import (
	"fmt"
	"os"
)

func main() {
	greeting()
}

func greeting() {
	hostname, error := os.Hostname()
	if error != nil {
		fmt.Println(error)
		return
	}
	username := os.Getenv("USERNAME")
	if username == "" {
		fmt.Println("Username not found")
		username = "Anonymous"
	}
	message := "Hello, " + username + ". I am running on " + hostname
	fmt.Println(message)
}
