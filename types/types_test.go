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
