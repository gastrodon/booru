package types

import (
	"strings"
	"testing"
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
