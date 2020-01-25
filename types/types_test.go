package types

import (
	"os"
	"testing"
	"time"
)

var test_live Client
var test_post Post
var test_user User
var test_pool Pool
var test_profile Profile
var now int64

func TestMain(main *testing.M) {
	test_live = Client{
		Host: "https://danbooru.donmai.us/",
	}
	test_post, _, _ = test_live.GetPost(1)
	test_user, _, _ = test_live.GetUser(581729)
	test_pool, _, _ = test_live.GetPool(2)
	test_profile, _, _ = test_auth.GetProfile()

	now = time.Now().Unix()
	os.Exit(main.Run())
}
