package main

import (
	"reflect"
	"testing"
)


func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name 	   string
		inputURL   string
		inputBody  string
		expected   []string
	} {
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
			name: "No <a> Tags",
			inputURL: "https://example.com/",
			inputBody: `
			<html>
				<body>
					<href="https://weather.com">Weather<>
					<href="https://weather.com">Weather<>
					<href="https://weather.com">Weather<>
					<href="https://weather.com">Weather<>
					<href="https://weather.com">Weather<>
					<href="https://weather.com">Weather<>
				</body>
			</html>
			`,
			expected: []string{},
		},
		{
			name: "Invalid HTML",
			inputURL: "https://example.com/",
			inputBody: `
			<html>
				<bod>
					<a href="https://weather.com">Weather</a>
					<a href="https://weather.com">Weather</a>
					<a href="https://weather.com">Weather</a>
				</bod>
			</html>
			`,
			expected: []string{"https://weather.com", "https://weather.com", "https://weather.com"},
		},
		{
			name: "Nested Tags",
			inputURL: "https://example.com/",
			inputBody: `
			<html>
				<body>
					<a href="https://weather.com">Weather
						<a href="https://example.com">Weather</a>
					</a>
					<a href="https://weather.com">Weather
						<a href="https://example.com">Weather</a>
					</a>
					<a href="https://weather.com">Weather
						<a href="https://example.com">Weather</a>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://weather.com", "https://example.com", "https://weather.com", "https://example.com", "https://weather.com", "https://example.com"},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return 
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}