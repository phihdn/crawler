// get_html.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// getHTML fetches the HTML content from the given URL
func getHTML(rawURL string) (string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check for error-level status code (400+)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error: %v", resp.Status)
	}

	// Check the Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content type %s is not text/html", contentType)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Return the HTML as a string
	return string(bodyBytes), nil
}
