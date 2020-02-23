package util

import (
	"testing"
	"time"
)

func Test_FormatQS(test *testing.T) {
	var expected string = "baz=qux&foo=bar"
	var qs map[string]string = map[string]string{
		"foo": "bar",
		"baz": "qux",
	}

	var formatted string = FormatQS(qs)
	if formatted != expected {
		test.Errorf("FormatQS result want %s, have %s", expected, formatted)
	}
}

func Test_FormatQS_None(test *testing.T) {
	var expected string = ""
	var qs map[string]string = map[string]string{}

	var formatted string = FormatQS(qs)
	if formatted != expected {
		test.Errorf("FormatQS result want %s, have %s", expected, formatted)
	}
}

func Test_FormatQS_Multiple(test *testing.T) {
	var expected string = "baz=qux&foo=bar&foo=buz&something=else"
	var qs map[string]string = map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	var more_qs map[string]string = map[string]string{
		"something": "else",
		"foo":       "buz",
	}

	var formatted string = FormatQS(qs, more_qs)
	if formatted != expected {
		test.Errorf("FormatQS result want %s, have %s", expected, formatted)
	}

}

func Test_TimeFromPtr(test *testing.T) {
	var long_ago string = "2007-08-16T17:55:36.905-04:00"

	var when *time.Time
	var err error
	when, err = TimeFromPtr(&long_ago)
	if err != nil {
		test.Fatal(err)
	}

	if when == nil {
		test.Errorf("time for %s (%p) is nil", long_ago, &long_ago)
		test.FailNow()
	}

	var stamp int64 = 1187301336
	if when.Unix() != stamp {
		test.Errorf("time mismatch! have: %d, want: %d", when.Unix(), stamp)
	}
}

func Test_TimeFromPtr_nil(test *testing.T) {
	var when *time.Time
	var err error
	when, err = TimeFromPtr(nil)
	if err != nil {
		test.Fatal(err)
	}

	if when != nil {
		test.Errorf("time for nil is not nil, but instead %d (%p)", when.Unix(), when)
	}
}

func Test_WrapQS(test *testing.T) {
	var expected_inverse map[string]string = map[string]string{
		"42069":  "foo[time]",
		"Mizuki": "foo[name]",
	}

	var base string = "foo"
	var qs map[string]string = map[string]string{
		"time": "42069",
		"name": "Mizuki",
	}

	var transformed map[string]string = WrapQS(base, qs)
	var key, value string
	for key, value = range transformed {
		if key != expected_inverse[value] {
			test.Errorf("wrapped key mismatch! have: %s, want: %s", key, expected_inverse[value])
		}
	}
}

func Test_CommonParams(test *testing.T) {
	var expected map[string]string = map[string]string{
		"limit":  "30",
		"page":   "1",
		"random": "true",
	}

	var query_strings map[string]string = CommonParams(1, 30, true)

	var key, value string
	for key, value = range query_strings {
		if expected[key] != value {
			test.Errorf("Value mismatch at %s! have: %s, want: %s", key, value, expected[value])
		}
	}
}

func Test_CommonParams_NoRandom(test *testing.T) {
	var query_strings map[string]string = CommonParams(1, 1, false)

	var exists bool
	_, exists = query_strings["random"]

	if exists {
		test.Errorf("Query string random should not exist but is %s", query_strings["random"])
	}
}
