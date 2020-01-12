package booru

import (
	"github.com/gastrodon/booru/types"

	"os"
	"testing"
)

var test_live types.Client

func TestMain(main *testing.M) {
	test_live = types.Client{
		Host: "https://testbooru.donmai.us/",
	}
	os.Exit(main.Run())
}

func Test_ClientAt(test *testing.T) {
	var host string = "foobar"
	var user types.Client = ClientAt(host)
	if user.Host != host {
		test.Errorf("user.Host have %s, want %s", user.Host, host)
	}
}
