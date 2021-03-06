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
	OkDate(test, test_post.CreatedAt, "test_post.CreatedAt")
}

func Test_UpdatedAt(test *testing.T) {
	OkDate(test, test_post.UpdatedAt, "test_post.UpdatedAt")

	var updated *time.Time
	var err error
	updated, err = test_post.UpdatedAt()
	if updated == nil {
		return
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
	OkDate(test, test_post.LastCommentAt, "test_post.LastCommentAt")
}

func Test_LastCommentBumpedAt(test *testing.T) {
	OkDate(test, test_post.LastCommentBumpedAt, "test_post.LastCommentBumpedAt")
}

func Test_LastNoteAt(test *testing.T) {
	OkDate(test, test_post.LastNoteAt, "test_post.LastNoteAt")
}

func Test_Uploader(test *testing.T) {
	OkUser(test, test_post.Uploader, test_post.UploaderID)
}

func Test_Approver(test *testing.T) {
	var results []Post
	var err error
	results, err = test_live.GetPosts([]string{"-approver:none"}, false, 1, 30, true)
	if err != nil {
		test.Fatal(err)
	}

	if len(results) == 0 {
		test.Error("No posts were retrieved")
	}

	var current Post
	var approver User
	for _, current = range results {
		OkUser(test, current.Approver, *current.ApproverID)
		approver, _, err = current.Approver()
		if approver.Client.Host != test_live.Host {
			test.Errorf("Host mismatch! want: %s, have: %s", test_live.Host, approver.Client.Host)
		}

	}
}

func Test_Approver_NotApproved(test *testing.T) {
	var results []Post
	var err error
	results, err = test_live.GetPosts([]string{"approver:none"}, false, 1, 30, true)
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
