package main

import (
	"net/url"
	"strings"
)

// normalizeURL takes a URL string and returns a normalized version of the URL
func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Convert the host to lowercase
	host := strings.ToLower(parsedURL.Host)

	// Remove port if present
	if colonIndex := strings.Index(host, ":"); colonIndex != -1 {
		host = host[:colonIndex]
	}

	// Get the path without trailing slash
	path := parsedURL.Path
	if len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	// Combine host and path to form normalized URL
	result := host + path

	return result, nil
}
