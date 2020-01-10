package util

import (
	"testing"
)

func Test_FormatQS(test *testing.T) {
	var expected string = "foo=bar&baz=qux"
	var qs map[string]string = map[string]string{
		"foo": "bar",
		"baz": "qux",
	}

	var formatted string = FormatQS(qs)
	if formatted != expected {
		test.Errorf("FormatQS result want %s, have %s", expected, formatted)
	}
}
