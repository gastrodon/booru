package types

import (
	"os"
	"testing"
)

var test_me Client
var test_post Post

func TestMain(main *testing.M) {
	test_me = Client{
		Host: "https://testbooru.donmai.us/",
	}
	test_post, _ = test_me.GetPost(1)
	os.Exit(main.Run())
}
