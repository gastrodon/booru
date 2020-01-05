package booru

import (
	"net/http"
	"os"
	"testing"
)

var test_me Client

func TestMain(main *testing.M) {
	test_me = Client{
		Host: "https://testbooru.donmai.us/",
	}
	os.Exit(main.Run())
}

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

func Test_ClientAt(test *testing.T) {
	var host string = "foobar"
	var user Client = ClientAt(host)
	if user.Host != host {
		test.Errorf("user.Host have %s, want %s", user.Host, host)
	}
}
