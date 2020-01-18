package types

import (
	"os"
	"testing"
)

var test_live Client
var test_post Post
var test_user User
var test_pool Pool

func TestMain(main *testing.M) {
	test_live = Client{
		Host: "https://danbooru.donmai.us/",
	}
	test_post, _, _ = test_live.GetPost(1)
	test_user, _, _ = test_live.GetUser(581729)
	test_pool, _, _ = test_live.GetPool(2)
	os.Exit(main.Run())
}
