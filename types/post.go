package types

import (
	"strings"
)

/*
 * Represents a media post on *booru
 */
type Post struct {
	Client Client
	ID     int
	// Tag info
	TagString          string `json:"tag_string"`
	ArtistTagString    string `json:"tag_string_artist"`
	CharacterTagString string `json:"tag_string_character"`
	CopyrightTagString string `json:"tag_string_copyright"`
	GeneralTagString   string `json:"tag_string_general"`
	MetaTagString      string `json:"tag_string_meta"`
	TagCount           int    `json:"tag_count"`
	ArtistTagCount     int    `json:"tag_count_artist"`
	CharacterTagCount  int    `json:"tag_count_character"`
	CopyrightTagCount  int    `json:"tag_count_copyright"`
	GeneralTagCount    int    `json:"tag_count_general"`
	MetaTagCount       int    `json:"tag_count_meta"`

	Banned       bool `json:"is_banned"`        // Are posts by this artist banned?
	Deleted      bool `json:"is_deleted"`       // Was this post deleted?
	Favorited    bool `json:"is_favorited"`     // Was this post favorited by the requesting client?
	Flagged      bool `json:"is_flagged"`       // Was this post flagged for moderator review by the requesting client?
	Pending      bool `json:"is_pending"`       // Is this post pending moderator approval?
	NoteLocked   bool `json:"is_note_locked"`   // Are the notes on this post locked?
	RatingLocked bool `json:"is_rating_locked"` // Is the rating on this post locked?
	StatusLocked bool `json:"is_status_locked"` // Is the status on this post locked?
}

/*
 * Get all of the tags of a type of this post
 *
 * artist: 		tags of the artists who made this post
 * character: 	tags of the characters in this post
 * copyright: 	tags of the source material for this post
 * general: 	tags that describe the content of this post
 * meta: 		tags that describe the properties of this post
 * all: 		all tags of all types
 */
func (post Post) Tags(tag_type string) (tags []string) {
	var splittable string
	switch tag_type {
	case "artist":
		splittable = post.ArtistTagString
		break
	case "character":
		splittable = post.CharacterTagString
		break
	case "copyright":
		splittable = post.CopyrightTagString
		break
	case "general":
		splittable = post.GeneralTagString
		break
	case "meta":
		splittable = post.MetaTagString
		break
	case "all":
	default:
		splittable = post.TagString
		break
	}

	tags = strings.Split(splittable, " ")
	return
}
