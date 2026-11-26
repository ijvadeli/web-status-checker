package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const arrow = "==>"

func main() {
	// String slice with urls
	websites := []string{
		"https://go.dev/learn",
		"https://www.w3schools.com/",
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

		// Get response status code from requested website
		response, err := http.Get(website)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

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
	}
}
