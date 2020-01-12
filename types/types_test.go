package types

import (
	"os"
	"testing"
)

var test_live Client
var test_post Post

func TestMain(main *testing.M) {
	test_live = Client{
		Host: "https://danbooru.donmai.us/",
	}
	test_post, _, _ = test_live.GetPost(1)
	os.Exit(main.Run())
}
