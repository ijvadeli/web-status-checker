package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const arrow = "==>"

func main() {
	websites := []string{
		"https://go.dev/",
		"https://www.w3schools.com/",
	}

	// Go through each website in websites slice
	for _, website := range websites {
		// Parse requested URI to validate if website exists
		u, err := url.ParseRequestURI(website)
		if err != nil {
			panic(err)
		}

		// Get response status code from requested website
		response, err := http.Get(website)
		if err != nil {
			panic(err)
		}

		fmt.Println("---------------------------------------")
		fmt.Println(arrow, u)
		fmt.Println(arrow, response.StatusCode, http.StatusText(response.StatusCode))

	}
}
