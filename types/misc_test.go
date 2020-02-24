package types

import (
	"testing"
)

func Test_ParseTags_RelatedTags(test *testing.T) {
	var related RelatedTagResponse
	var err error
	related, err = test_live.GetRelatedTags("solo", "all")
	if err != nil {
		test.Fatal(err)
	}

	if len(related.RelatedTags) != 0 {
		test.Errorf("related.RelatedTags is already populated! have: %v", related.RelatedTags)
	}

	related.ParseTags()
	if len(related.RelatedTags) == 0 {
		test.Error("related.RelatedTags was not populated!")
	}
}

func Test_ParseTags_RelatedWikiTags(test *testing.T) {
	var related RelatedTagResponse
	var err error
	related, err = test_live.GetRelatedTags("solo", "all")
	if err != nil {
		test.Fatal(err)
	}

	if len(related.RelatedWikiTags) != 0 {
		test.Errorf("related.RelatedWikiTags is already populated! have: %v", related.RelatedWikiTags)
	}

	related.ParseWikiTags()
	if len(related.RelatedWikiTags) == 0 {
		test.Error("related.RelatedWikiTags was not populated!")
	}
}
