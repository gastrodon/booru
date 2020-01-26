package types

import (
	"testing"
	"time"
)

func Test_CreatedAt_Pool(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_pool.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix()-1000 >= now {
		test.Errorf("pool.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_UpdatedAt_Pool(test *testing.T) {
	var updated *time.Time
	var err error
	updated, err = test_pool.UpdatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated == nil {
		return
	}

	if updated.Unix() >= now {
		test.Errorf("pool.UpdatedAt is in the future: %d", updated.Unix())
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

func Test_Creator(test *testing.T) {
	var creator User
	var exists bool
	var err error
	creator, exists, err = test_pool.Creator()
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("user %d does not exist", test_pool.CreatorID)
	}

	if creator.ID != test_pool.CreatorID {
		test.Errorf("creator ID mismatch! have: %d, want: %d", creator.ID, test_pool.CreatorID)
	}

	if creator.Name != test_pool.CreatorName {
		test.Errorf("creator name mismatch! have: %s, want: %s", creator.Name, test_pool.CreatorName)
	}
}
