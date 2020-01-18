package util

import (
	"fmt"
	"sort"
	"strings"
	"time"
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

	sort.Strings(parts)
	return strings.Join(parts, "&")
}

func TimeFromPtr(time_string *string) (parsed *time.Time, err error) {
	var _time time.Time
	if time_string != nil {
		_time, err = time.Parse(time.RFC3339, *time_string)
		parsed = &_time
	}
	return
}
