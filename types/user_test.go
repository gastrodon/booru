package types

import (
	"testing"
)

func Test_CreatedAt_User(test *testing.T) {
	OkDate(test, test_user.CreatedAt, "test_user.CreatedAt")
}
