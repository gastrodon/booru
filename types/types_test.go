package types

import (
	"os"
	"testing"
	"time"
)

var login, key string

var test_live, test_auth Client
var test_post Post
var test_user User
var test_pool Pool
var test_comment Comment
var test_profile Profile
var now int64

func TestMain(main *testing.M) {
	login = os.Getenv("BOORU_LOGIN")
	key = os.Getenv("BOORU_KEY")

	test_live = Client{
		Host: "https://danbooru.donmai.us/",
	}

	test_auth = Client{
		Host: "https://danbooru.donmai.us/",
	}
	test_auth.Auth(login, key)

	test_post, _, _ = test_live.GetPost(1)
	test_user, _, _ = test_live.GetUser(581729)
	test_pool, _, _ = test_live.GetPool(2)
	test_comment, _, _ = test_live.GetComment(1)
	test_profile, _, _ = test_auth.GetProfile()

	now = time.Now().Unix()
	os.Exit(main.Run())
}

func OkDate(test *testing.T, callable func() (*time.Time, error), label string) {
	var stamp *time.Time
	var err error
	stamp, err = callable()
	if err != nil {
		test.Fatal(err)
	}

	if stamp == nil {
		return
	}

	if stamp.Unix()-1000 >= now {
		test.Errorf("%s is in the future: %d", label, stamp.Unix())
	}
}

func OkUser(test *testing.T, callable func() (User, bool, error), id int) {
	var user User
	var exists bool
	var err error
	user, exists, err = callable()
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("User %d does not exist!", id)
	}

	if user.ID != id {
		test.Errorf("User id mismatch! have: %d, want: %d", user.ID, id)
	}
}

func OkPost(test *testing.T, callable func() (Post, bool, error), id int) {
	var post Post
	var exists bool
	var err error
	post, exists, err = callable()
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("#%d does not exist!", id)
	}

	if post.ID != id {
		test.Errorf("Post id mismatch! have: %d, want: %d", post.ID, id)
	}
}

func get_callable_user(id int) (callable func() (User, bool, error)) {
	callable = func() (User, bool, error) {
		return test_live.GetUser(id)
	}
	return
}

func get_callable_post(id int) (callable func() (Post, bool, error)) {
	callable = func() (Post, bool, error) {
		return test_live.GetPost(id)
	}
	return
}
