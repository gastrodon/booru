package types

import (
	"testing"
	"time"
)

func Test_CreatedAt_User(test *testing.T) {
	var now int64 = time.Now().Unix()

	var stamp *time.Time
	var err error
	stamp, err = test_user.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() >= now {
		test.Errorf("post.CreatedAt is in the future: %d", stamp.Unix())
	}
}
