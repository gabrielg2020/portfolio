package utils

import "strings"

// ParseStringArray takes a comma-delimited string and returns a slice of strings.
//
// Example:
//
//	s := "hello,world"
//	result := ParseStringArray(s) // ["hello", "world"]
func ParseStringSlice(s string) []string {
	if s == "" {
		return []string{}
	}

	// Split by comma and trim spaces
	items := strings.Split(s, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}

	return items
}