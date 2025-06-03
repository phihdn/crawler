package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Skip the program name (os.Args[0])

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// At this point, we have exactly one argument
	baseURL := args[0]
	fmt.Printf("starting crawl of: %s\n", baseURL)

	// Initialize the pages map to track the count of each URL
	pages := make(map[string]int)

	// Start crawling from the base URL
	crawlPage(baseURL, baseURL, pages)

	// Print the results
	fmt.Println("\nCrawl results:")
	fmt.Println("==============")
	for url, count := range pages {
		fmt.Printf("%s: %d\n", url, count)
	}
}
