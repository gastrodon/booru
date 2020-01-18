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
	response, err = test_live.make_request(method, where, qs_map, nil)

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
	var exists bool
	var err error
	post, exists, err = test_live.GetPost(id)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("#%d does not exist", id)
	}

	if post.ID != id {
		test.Errorf("ID mismatch have: %d, want: %d", post.ID, id)
	}
}

func Test_GetPost_NoSuchPost(test *testing.T) {
	var id = -1

	var exists bool
	var err error
	_, exists, err = test_live.GetPost(id)
	if err != nil {
		test.Fatal(err)
	}

	if exists {
		test.Errorf("#%d exists", id)
	}
}

func Test_GetUser(test *testing.T) {
	var id int = 9

	var user User
	var exists bool
	var err error
	user, exists, err = test_live.GetUser(id)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("user %d does not exist", id)
	}

	if user.ID != id {
		test.Errorf("ID mismatch have: %d, want: %d", user.ID, id)
	}
}

func Test_GetUser_NoSuchUser(test *testing.T) {
	var id = -1

	var exists bool
	var err error
	_, exists, err = test_live.GetUser(id)
	if err != nil {
		test.Fatal(err)
	}

	if exists {
		test.Errorf("user %d exists", id)
	}
}

func Test_GetPosts(test *testing.T) {
	var tags []string = []string{"solo"}
	var page int = 1
	var limit int = 20

	var results []Post
	var err error
	results, err = test_live.GetPosts(tags, page, limit, false, false)
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
	results, err = test_live.GetPosts([]string{}, page, 1, false, false)
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
	results, err = test_live.GetPosts([]string{}, 1, 100, true, false)
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
	var exists bool
	var err error
	post, exists, err = test_live.GetPostMD5(md5)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("Post #%d with md5 %s does not exist", post.ID, post.MD5)
	}

	if post.ID != test_post.ID {
		test.Errorf("Post id mismatch: have %d, want %d", post.ID, test_post.ID)
	}
}

func Test_GetPostMD5_NoSuchPost(test *testing.T) {
	var exists bool
	var err error
	_, exists, err = test_live.GetPostMD5("_")
	if err != nil {
		test.Fatal(err)
	}

	if exists {
		test.Error("Post with md5 _ exists")
	}
}

func Test_GetPool(test *testing.T) {
	var id int = 2

	var pool Pool
	var exists bool
	var err error
	pool, exists, err = test_live.GetPool(id)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("pool %d does not exist", id)
	}

	if len(pool.PostIDs) != pool.PostCount {
		test.Errorf("pool post count mismatch! pool.PostIDs len: %d, pool.PostCount: %d", len(pool.PostIDs), pool.PostCount)
	}
}
