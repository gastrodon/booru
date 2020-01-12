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

func Test_GetPost(test *testing.T) {
	var id int = 2

	var post Post
	var err error
	post, err = test_me.GetPost(id)
	if err != nil {
		test.Fatal(err)
	}

	if post.ID != id {
		test.Errorf("ID mismatch have: %d, want: %d", post.ID, id)
	}
}

func Test_GetUser(test *testing.T) {
	var id int = 9

	var user User
	var err error
	user, err = test_me.GetUser(id)
	if err != nil {
		test.Fatal(err)
	}

	if user.ID != id {
		test.Errorf("ID mismatch have: %d, want: %d", user.ID, id)
	}
}
