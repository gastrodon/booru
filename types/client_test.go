package types

import (
	"net/http"
	"testing"
)

func Test_method_make_request(test *testing.T) {
	var method = "GET"
	var where = "/posts"
	var qs_map map[string]string = map[string]string{
		"random": "true",
	}

	var response *http.Response
	var err error
	response, err = test_me.make_request(method, where, qs_map, nil)

	if err != nil {
		test.Fatal(err)
	}

	if response.StatusCode != 200 {
		test.Errorf("response.StatusCode is %d", response.StatusCode)
	}
}
