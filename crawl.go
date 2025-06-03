package main

import (
	"fmt"
	"net/url"
)

// crawlPage recursively crawls a website starting from the current URL
func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Make sure the raw current URL is on the same domain as the raw base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("failed to parse base URL: %v\n", err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse current URL: %v\n", err)
		return
	}

	// Check if the current URL is on the same domain as the base URL
	if baseURL.Host != currentURL.Host {
		fmt.Printf("skipping %s: not on same domain as %s\n", rawCurrentURL, rawBaseURL)
		return
	}

	// Normalize the current URL
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// If we've already crawled this page, just increment the count and return
	if count, ok := pages[normalizedURL]; ok {
		pages[normalizedURL] = count + 1
		return
	}

	// Otherwise, add an entry to the pages map for the normalized URL with count 1
	pages[normalizedURL] = 1

	// Print a message to show progress
	fmt.Printf("crawling: %s\n", rawCurrentURL)

	// Get the HTML from the current URL
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to get HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Extract all URLs from the HTML
	urls, err := getURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to get URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Recursively crawl each URL found on the page
	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
