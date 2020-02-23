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
