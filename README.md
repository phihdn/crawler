# Web Crawler

A simple web crawler that crawls a website, counts links between pages, and generates a report on internal link structure.

## Description

This Go application crawls a specified website's internal pages and counts how many times each page is linked from other pages on the site. It helps analyze the internal link structure of a website.

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.18 or higher)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/phihdn/crawler.git
   cd crawler
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

### Usage

Build and run the crawler:

```bash
go build -o crawler
./crawler [starting URL] [max concurrency] [max pages]
```

Parameters:

- `starting URL`: The URL to start crawling from
- `max concurrency`: Maximum number of concurrent HTTP requests
- `max pages`: Maximum number of pages to crawl

Example:

```bash
./crawler https://example.com 10 100
```

This will crawl example.com with up to 10 concurrent requests and a maximum of 100 pages.

### Output Example

When the crawler completes its task, it will generate a report like this:

```text
=============================
  REPORT for https://example.com
=============================
Found 42 internal links to https://example.com/page1
Found 36 internal links to https://example.com/about
Found 28 internal links to https://example.com/contact
Found 15 internal links to https://example.com/blog/post1
...
```

The report shows each page sorted by the number of internal links pointing to it (in descending order).

## Project Structure

- `main.go` - Entry point and report generation
- `crawl.go` - Core crawling functionality
- `get_html.go` - HTML fetching functions
- `get_urls_from_html.go` - URL extraction from HTML content
- `normalize_url.go` - URL normalization functions
- `configure.go` - Configuration setup

## Ideas for Extending the Project

- Make the script run on a timer and deploy it to a server. Have it email you every so often with a report.
- Add more robust error checking so that you can crawl larger sites without issues.
- Count external links, as well as internal links, and add them to the report.
- Save the report as a CSV spreadsheet rather than printing it to the console.
- Use a graphics library to create an image that shows the links between the pages as a graph visualization.
- Make requests concurrently to speed up the crawling process.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Created as part of a [Boot.dev](https://boot.dev) guided project.
