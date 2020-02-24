package types

import (
	"testing"
	"time"
)

func Test_CreatedAt_Comment(test *testing.T) {
	OkDate(test, test_comment.CreatedAt, "test_comment.CreatedAt")
}

func Test_UpdatedAt_Comment(test *testing.T) {
	OkDate(test, test_comment.UpdatedAt, "test_comment.UpdatedAt")

	var updated *time.Time
	var err error
	updated, err = test_comment.UpdatedAt()
	if updated == nil {
		return
	}

	var created *time.Time
	created, err = test_comment.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() < created.Unix() {
		test.Errorf("test_comment.UpdatedAt is before test_comment.CreatedAt: %d < %d", updated.Unix(), created.Unix())
	}
}

func Test_GetPost_Comment(test *testing.T) {
	OkPost(test, test_comment.GetPost, test_comment.PostID)
}

func Test_GetCreator_Comment(test *testing.T) {
	OkUser(test, test_comment.GetCreator, test_comment.CreatorID)
}

func Test_GetUpdater_Comment(test *testing.T) {
	OkUser(test, test_comment.GetUpdater, test_comment.UpdaterID)
}
