package util

import (
	"fmt"
	"strings"
)

/*
 * Format a map of query strings into a string that may be part of a request
 */
func FormatQS(query_strings map[string]string) (formatted string) {
	var parts []string
	var key, value string
	for key, value = range query_strings {
		parts = append(parts, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(parts, "&")
}
