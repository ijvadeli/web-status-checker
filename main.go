package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const arrow = "==>"

func main() {
	// String slice with urls
	websites := []string{
		"https://go.dev",
		"https://github.com/ijvadeli/web-status-checker",
	}

	// Directory for saving crawled files
	dir := "crawled"

	// Go through each website in websites slice
	for _, website := range websites {
		// Parse requested URI to validate if website exists
		u, err := url.ParseRequestURI(website)
		if err != nil {
			panic(err)
		}

		// Limiting requests and creating global client
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		// Start timer
		timer := time.Now()

		// Create a HTTP request
		response, err := client.Get(website)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		responseTime := time.Since(timer)

		// Crawl website and read response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		// Create crawled site file and save to crawled
		filename := u.Hostname() + ".html"
		file, err := os.Create(filepath.Join(dir, filepath.Base(filename)))
		if err != nil {
			panic(err)
		}
		file.Write(body)

		fmt.Println("---------------------------------------")
		fmt.Println(arrow, u.Scheme+"://"+u.Host)
		fmt.Println(arrow, response.StatusCode, http.StatusText(response.StatusCode))
		fmt.Println(arrow, responseTime)
	}
}
