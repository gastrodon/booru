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

func compare_tags(test *testing.T, parsed []string, raw string) {
	var joined string = strings.Join(parsed, " ")
	if raw != joined {
		test.Errorf("test_post.TagString mismatch raw: %s, joined: %s", raw, joined)
	}
}

func Test_tag_types(test *testing.T) {
	compare_tags(test, test_post.Tags(""), test_post.TagString)
	compare_tags(test, test_post.Tags("artist"), test_post.ArtistTagString)
	compare_tags(test, test_post.Tags("character"), test_post.CharacterTagString)
	compare_tags(test, test_post.Tags("copyright"), test_post.CopyrightTagString)
	compare_tags(test, test_post.Tags("general"), test_post.GeneralTagString)
	compare_tags(test, test_post.Tags("meta"), test_post.MetaTagString)
}

func Test_CreatedAt(test *testing.T) {
	var now int64 = time.Now().Unix()

	var created time.Time
	var err error
	created, err = test_post.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if created.Unix() >= now {
		test.Errorf("post.CreatedAt is in the future: %d", created.Unix())
	}
}

func Test_UpdatedAt(test *testing.T) {
	var now int64 = time.Now().Unix()

	var updated time.Time
	var err error
	updated, err = test_post.UpdatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() >= now {
		test.Errorf("post.UpdatedAt is in the future: %d", updated.Unix())
	}

	var created time.Time
	created, err = test_post.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() < created.Unix() {
		test.Errorf("post.UpdatedAt is before post.CreatedAt: %d < %d", updated.Unix(), created.Unix())
	}
}
