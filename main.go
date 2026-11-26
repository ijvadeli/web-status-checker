package main

import (
	"fmt"
	"net/url"
)

func main() {
	websites := []string{
		"https://go.dev/",
		"https://w3schools.com",
	}

	for _, website := range websites {
		u, err := url.ParseRequestURI(website)
		if err != nil {
			panic(err)
		}

		fmt.Println("URL => ", u)
	}
}
