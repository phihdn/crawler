package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "multiple relative URLs",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a href="/products">Products</a>
		<a href="/about">About Us</a>
		<a href="/contact">Contact</a>
		<div>
			<a href="/blog">Blog</a>
		</div>
	</body>
</html>
`,
			expected: []string{
				"https://example.com/products",
				"https://example.com/about",
				"https://example.com/contact",
				"https://example.com/blog",
			},
		},
		{
			name:     "nested URLs and fragment identifiers",
			inputURL: "https://test.org",
			inputBody: `
<html>
	<body>
		<div>
			<a href="https://external.org/path">External</a>
			<ul>
				<li><a href="/relative/path">Relative</a></li>
				<li><a href="#section">Fragment</a></li>
				<li><a href="/path/with#fragment">Path with fragment</a></li>
			</ul>
		</div>
	</body>
</html>
`,
			expected: []string{
				"https://external.org/path",
				"https://test.org/relative/path",
				"https://test.org#section",
				"https://test.org/path/with#fragment",
			},
		},
		{
			name:     "URLs with query parameters",
			inputURL: "https://search.com",
			inputBody: `
<html>
	<body>
		<a href="/search?q=golang">Search Golang</a>
		<a href="https://api.search.com/v1/search?q=rust&lang=en">Search Rust</a>
		<a href="/filters?category=programming&level=advanced">Advanced Programming</a>
	</body>
</html>
`,
			expected: []string{
				"https://search.com/search?q=golang",
				"https://api.search.com/v1/search?q=rust&lang=en",
				"https://search.com/filters?category=programming&level=advanced",
			},
		},
		{
			name:     "empty href values and malformed URLs",
			inputURL: "https://blog.example.com",
			inputBody: `
<html>
	<body>
		<a href="">Home</a>
		<a>No href attribute</a>
		<a href="   ">Whitespace href</a>
		<a href="/valid/path">Valid path</a>
	</body>
</html>
`,
			expected: []string{
				"https://blog.example.com", // For empty href
				"https://blog.example.com", // For whitespace href
				"https://blog.example.com/valid/path",
			},
		},
		{
			name:     "URLs with different schemes",
			inputURL: "http://legacy.site.net",
			inputBody: `
<html>
	<body>
		<a href="https://secure.site.net">Secure</a>
		<a href="http://old.site.net">Old</a>
		<a href="mailto:contact@site.net">Email us</a>
		<a href="tel:+123456789">Call us</a>
		<a href="ftp://files.site.net">Downloads</a>
	</body>
</html>
`,
			expected: []string{
				"https://secure.site.net",
				"http://old.site.net",
				"mailto:contact@site.net",
				"tel:+123456789",
				"ftp://files.site.net",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("getURLsFromHTML(%s, %s) returned unexpected error: %v", tc.inputBody, tc.inputURL, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("getURLsFromHTML(%s, %s) = %v, expected %v", tc.inputBody, tc.inputURL, actual, tc.expected)
			}
		})
	}
}

func TestGetURLsFromHTMLErrors(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
	}{
		// Note: HTML parsers are very forgiving, so it's difficult to create HTML that won't parse.
		// Instead, we'll focus on testing the base URL parsing error
		{
			name:      "invalid base URL",
			inputURL:  "://invalid-url",
			inputBody: "<html><body><a href='/path'>Test</a></body></html>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err == nil {
				t.Errorf("getURLsFromHTML(%s, %s) did not return an error as expected", tc.inputBody, tc.inputURL)
			}
		})
	}
}
