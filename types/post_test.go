package types

import (
	"strings"
	"testing"
	"time"
)

func Test_Tags(test *testing.T) {
	var tags []string = test_post.Tags("")
	if len(tags) == 0 {
		test.Errorf("test_post.Tags from testbooru are empty")
	}

	var general []string = test_post.Tags("general")
	if len(general) >= len(tags) {
		test.Errorf("test_post.Tags there are more general tags than all tags")
	}
}

func compareTags(test *testing.T, parsed []string, raw string) {
	var joined string = strings.Join(parsed, " ")
	if raw != joined {
		test.Errorf("test_post.TagString mismatch raw: %s, joined: %s", raw, joined)
	}
}

func Test_tag_types(test *testing.T) {
	compareTags(test, test_post.Tags(""), test_post.TagString)
	compareTags(test, test_post.Tags("all"), test_post.TagString)
	compareTags(test, test_post.Tags("artist"), test_post.ArtistTagString)
	compareTags(test, test_post.Tags("character"), test_post.CharacterTagString)
	compareTags(test, test_post.Tags("copyright"), test_post.CopyrightTagString)
	compareTags(test, test_post.Tags("general"), test_post.GeneralTagString)
	compareTags(test, test_post.Tags("meta"), test_post.MetaTagString)
}

func Test_CreatedAt(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_post.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() - 1000 >= now {
		test.Errorf("post.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_UpdatedAt(test *testing.T) {
	var updated *time.Time
	var err error
	updated, err = test_post.UpdatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated == nil {
		return
	}

	if updated.Unix() >= now {
		test.Errorf("post.UpdatedAt is in the future: %d", updated.Unix())
	}

	var created *time.Time
	created, err = test_post.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() < created.Unix() {
		test.Errorf("post.UpdatedAt is before post.CreatedAt: %d < %d", updated.Unix(), created.Unix())
	}
}

func Test_LastCommentAt(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_post.LastCommentAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp == nil {
		return
	}

	if stamp.Unix() - 1000 >= now {
		test.Errorf("post.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_LastCommentBumpedAt(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_post.LastCommentBumpedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp == nil {
		return
	}

	if stamp.Unix() - 1000 >= now {
		test.Errorf("post.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_LastNoteAt(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_post.LastNoteAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp == nil {
		return
	}

	if stamp.Unix() - 1000 >= now {
		test.Errorf("post.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_Uploader(test *testing.T) {
	var uploader User
	var exists bool
	var err error
	uploader, exists, err = test_post.Uploader()
	if err != nil {
		test.Fatal(err)
	}

	if uploader.Client.Host != test_live.Host {
		test.Errorf("Host mismatch! want: %s, have: %s", test_live.Host, uploader.Client.Host)
	}

	if !exists {
		test.Errorf("#%d uploader %d does not exist", test_post.ID, test_post.UploaderID)
	}

	if uploader.ID != test_post.UploaderID {
		test.Errorf("Uploader id mismatch! have: %d, want: %d", uploader.ID, test_post.UploaderID)
	}
}

func Test_Approver(test *testing.T) {
	var results []Post
	var err error
	results, err = test_live.GetPosts([]string{"-approver:none"}, 1, 30, true, false)
	if err != nil {
		test.Fatal(err)
	}

	if len(results) == 0 {
		test.Error("No posts were retrieved")
	}

	var current Post
	var approver User
	var exists bool
	for _, current = range results {
		approver, exists, err = current.Approver()
		if err != nil {
			test.Fatal(err)
		}

		if current.ApproverID == nil {
			test.Errorf("#%d is not approved", current.ID)
			break
		}

		if approver.Client.Host != test_live.Host {
			test.Errorf("Host mismatch! want: %s, have: %s", test_live.Host, approver.Client.Host)
		}

		if !exists {
			test.Errorf("#%d approver %d does not exist", current.ID, *current.ApproverID)
		}

		if approver.ID != *current.ApproverID {
			test.Errorf("Uploader id mismatch on #%d! have: %d, want: %d", current.ID, approver.ID, *current.ApproverID)
		}

	}
}

func Test_Approver_NotApproved(test *testing.T) {
	var results []Post
	var err error
	results, err = test_live.GetPosts([]string{"approver:none"}, 1, 30, true, false)
	if err != nil {
		test.Fatal(err)
	}

	if len(results) == 0 {
		test.Error("No posts were retrieved")
	}

	var current Post
	var approver User
	var exists bool
	for _, current = range results {
		approver, exists, err = current.Approver()
		if err != nil {
			test.Fatal(err)
		}

		if current.ApproverID != nil {
			test.Errorf("#%d is approved by %d", current.ID, *current.ApproverID)
		}

		if exists {
			test.Errorf("unapproved #%d approver is user %d", current.ID, approver.ID)
		}
	}

}
