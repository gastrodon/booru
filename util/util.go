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
func FormatQS(query_strings ...map[string]string) (formatted string) {
	var total_size int = 0
	var current map[string]string
	for _, current = range query_strings {
		total_size += len(current)
	}

	var parts []string = make([]string, total_size)

	var index int = 0
	var key, value string
	for _, current = range query_strings {
		for key, value = range current {
			parts[index] = fmt.Sprintf("%s=%s", key, value)
			index += 1
		}
	}

	sort.Strings(parts)
	return strings.Join(parts, "&")
}

/*
 * Given a base string, transform a map of query string such that for each string
 * key: value -> base[string]: value
 *
 * This is used to simplify making query string searches on danbooru
 */
func WrapQS(base string, query_strings map[string]string) (formatted map[string]string) {
	formatted = map[string]string{}
	base = fmt.Sprintf("%s[%%s]", base)

	var key, value string
	for key, value = range query_strings {
		formatted[fmt.Sprintf(base, key)] = value
	}

	return
}

func TimeFromPtr(time_string *string) (parsed *time.Time, err error) {
	var _time time.Time
	if time_string != nil {
		_time, err = time.Parse(time.RFC3339, *time_string)
		parsed = &_time
	}
	return
}
