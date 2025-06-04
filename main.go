package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type pageCount struct {
	url   string
	count int
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("maxConcurrency must be a positive integer")
		return
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil || maxPages < 1 {
		fmt.Println("maxPages must be a positive integer")
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")

	// Sort the pages by count (desc) and URL (asc)
	sortedPages := sortPages(pages)

	// Print the report
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
	}
}

func sortPages(pages map[string]int) []pageCount {
	// Convert the map to a slice of structs
	var pageSlice []pageCount
	for url, count := range pages {
		pageSlice = append(pageSlice, pageCount{url: url, count: count})
	}

	// Sort by count (desc) and URL (asc)
	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].count == pageSlice[j].count {
			return pageSlice[i].url < pageSlice[j].url // Alphabetical when counts are equal
		}
		return pageSlice[i].count > pageSlice[j].count // Descending order by count
	})

	return pageSlice
}
