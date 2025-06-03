package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// getURLsFromHTML extracts all URLs from anchor tags in HTML and converts relative URLs to absolute URLs
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	reader := strings.NewReader(htmlBody)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	extractURLs(doc, parsedBaseURL, &urls)

	return urls, nil
}

// extractURLs recursively traverses the HTML node tree and extracts URLs from anchor tags
func extractURLs(n *html.Node, baseURL *url.URL, urls *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		found := false
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				// Skip empty or whitespace-only href values
				if strings.TrimSpace(attr.Val) == "" {
					// If href is empty, use the base URL
					*urls = append(*urls, baseURL.String())
					found = true
					break
				}

				// Parse the URL found in the href attribute
				parsedURL, err := url.Parse(attr.Val)
				if err != nil {
					// Skip malformed URLs
					continue
				}

				// Resolve the URL against the base URL (to convert relative URLs to absolute)
				resolvedURL := baseURL.ResolveReference(parsedURL)
				*urls = append(*urls, resolvedURL.String())
				found = true
				break
			}
		}

		// Skip nodes that don't have an href attribute or have a whitespace-only href
		if !found {
			// No need to add any URL
		}
	}

	// Recursively process child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractURLs(c, baseURL, urls)
	}
}
