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

func Test_GetPost_NoSuchPost(test *testing.T) {
	var id = -1

	var post Post
	var err error
	post, err = test_me.GetPost(id)
	if err == nil {
		test.Error("No error is returned")
	}

	if post.ID != 0 {
		test.Errorf("#%d was retrieved", post.ID)
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

func Test_GetUser_NoSuchUser(test *testing.T) {
	var id = -1

	var user User
	var err error
	user, err = test_me.GetUser(id)
	if err == nil {
		test.Error("No error is returned")
	}

	if user.ID != 0 {
		test.Errorf("user %d was retrieved", user.ID)
	}
}

func Test_GetPosts(test *testing.T) {
	var tags []string = []string{"solo"}
	var page int = 1
	var limit int = 20

	var results []Post
	var err error
	results, err = test_me.GetPosts(tags, page, limit, false, false)
	if err != nil {
		test.Fatal(err)
	}

	if len(results) > limit {
		test.Errorf("Too many results were returned! have: %d, want: >= %d", len(results), limit)
	}

	var current Post
	var current_tags []string
	var index int
	var tag string
	for _, current = range results {

		current_tags = current.Tags("")
		for index, tag = range current_tags {
			if tag == tags[0] {
				break
			} else if index+1 == len(current_tags) {
				test.Errorf("Post #%d does not have the tag %s (index %d)", current.ID, tags[0], len(tags))
			}
		}
	}
}

func Test_GetPosts_TooManyPages(test *testing.T) {
	var page int = 10000

	var results []Post
	var err error
	results, err = test_me.GetPosts([]string{}, page, 1, false, false)
	if err == nil {
		test.Error("No error is returned")
	}

	if len(results) != 0 {
		test.Errorf("%d posts were returned, starting with #%d", len(results), results[0].ID)
	}
}

func Test_GetPosts_Random(test *testing.T) {
	var results []Post
	var err error
	results, err = test_me.GetPosts([]string{}, 1, 100, true, false)
	if err != nil {
		test.Fatal(err)
	}

	var cap int = len(results) - 1
	var head int = 1
	for head != cap {
		if results[head].ID != results[head-1].ID+1 {
			return
		}
		head += 1
	}

	test.Errorf("All posts were sequential")
}

func Test_GetPostMD5(test *testing.T) {
	var md5 string = test_post.MD5

	var post Post
	var err error
	post, err = test_me.GetPostMD5(md5)
	if err != nil {
		test.Fatal(err)
	}

	if post.ID != test_post.ID {
		test.Errorf("Post id mismatch: have %d, want %d", post.ID, test_post.ID)
	}
}

func Test_GetPostMD5_NoSuchPost(test *testing.T) {
	var post Post
	var err error
	post, err = test_me.GetPostMD5("")
	if err == nil {
		test.Error("No error is returned")
	}

	if post.ID != 0 {
		test.Errorf("#%d was retrieved", post.ID)
	}
}

func Test_PostMD5Exists(test *testing.T) {
	var exists bool
	var err error
	exists, err = test_me.PostMD5Exists(test_post.MD5)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("Post #%d with md5 %s returns as not existing", test_post.ID, test_post.MD5)
	}
}

func Test_PostMD5Exists_NoSuchPost(test *testing.T) {
	var exists bool
	var err error
	exists, err = test_me.PostMD5Exists("_")
	if err != nil {
		test.Fatal(err)
	}

	if exists {
		test.Errorf("Empty md5 exists")
	}
}
