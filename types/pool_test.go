package types

import (
	"testing"
	"time"
)

func Test_CreatedAt_Pool(test *testing.T) {
	OkDate(test, test_pool.CreatedAt, "pool.CreatedAt")
}

func Test_UpdatedAt_Pool(test *testing.T) {
	OkDate(test, test_pool.UpdatedAt, "pool.UpdatedAt")

	var updated *time.Time
	var err error
	updated, err = test_pool.UpdatedAt()
	if updated == nil {
		return
	}

	var created *time.Time
	created, err = test_pool.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() < created.Unix() {
		test.Errorf("pool.UpdatedAt is before pool.CreatedAt: %d < %d", updated.Unix(), created.Unix())
	}
}

func Test_PostsRange(test *testing.T) {
	var start int = 2
	var stop int = test_pool.PostCount - 2
	var diff int = 4

	var posts []Post
	var err error
	posts, err = test_pool.PostsRange(start, stop)
	if err != nil {
		test.Fatal(err)
	}

	if len(posts) != test_pool.PostCount-diff {
		test.Errorf("Post count mismatch! have: %d, want: %d", len(posts), test_pool.PostCount-diff)
	}

	return
	var index int
	var current Post
	for index, current = range posts[start:stop] {
		if test_pool.PostIDs[index] != current.ID {
			test.Errorf("Post mismatch at %d! have: #%d, want: #%d", index, current.ID, test_pool.PostIDs[index])
		}

		if current.Client.Host != test_pool.Client.Host {
			test.Errorf("Post client mismatch! %d does not have a client at %s", index, test_pool.Client.Host)
		}
	}
}

func Test_Posts(test *testing.T) {
	var posts []Post
	var err error
	posts, err = test_pool.Posts()
	if err != nil {
		test.Fatal(err)
	}

	if len(posts) != len(test_pool.PostIDs) {
		test.Errorf("Post count mismatch! have: %d, want: %d", len(posts), len(test_pool.PostIDs))
	}

	var index int
	var current Post
	for index, current = range posts {
		if test_pool.PostIDs[index] != current.ID {
			test.Errorf("Post mismatch at %d! have: #%d, want: #%d", index, current.ID, test_pool.PostIDs[index])
		}

		if current.Client.Host != test_pool.Client.Host {
			test.Errorf("Post client mismatch! %d does not have a client at %s", index, test_pool.Client.Host)
		}
	}
}

func Test_PostAt(test *testing.T) {
	var index int = 2
	var id int = test_pool.PostIDs[index]

	var expected Post
	var exists bool
	var err error
	expected, exists, err = test_live.GetPost(id)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("Post #%d (test_pool post %d) does not exist", id, index)
		test.FailNow()
	}

	var post Post
	post, exists, err = test_pool.PostAt(index)
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("Post #%d (test_pool post %d) does not exist", id, index)
		test.FailNow()
	}

	if post.ID != expected.ID {
		test.Errorf("ID mismatch at %d! have: #%d, want: #%d", index, post.ID, expected.ID)
	}
}
