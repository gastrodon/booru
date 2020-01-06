package booru

import (
	"strings"
)

/*
 * Represents a media post on *booru
 */
type Post struct {
	Client Client
	// Tag info
	tag_string           string `json:"tag_string"`
	tag_string_artist    string `json:"tag_string_artist"`
	tag_string_character string `json:"tag_string_character"`
	tag_string_copyright string `json:"tag_string_copyright"`
	tag_string_general   string `json:"tag_string_general"`
	tag_string_meta      string `json:"tag_string_meta"`
	tag_count            uint   `json:"tag_count"`
	tag_count_artist     uint   `json:"tag_count_artist"`
	tag_count_character  uint   `json:"tag_count_character"`
	tag_count_copyright  uint   `json:"tag_count_copyright"`
	tag_count_general    uint   `json:"tag_count_general"`
	tag_count_meta       uint   `json:"tag_count_meta"`

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
		splittable = post.tag_string_artist
		break
	case "character":
		splittable = post.tag_string_character
		break
	case "copyright":
		splittable = post.tag_string_copyright
		break
	case "general":
		splittable = post.tag_string_character
		break
	case "meta":
		splittable = post.tag_string_meta
		break
	case "all":
	default:
		splittable = post.tag_string
		break
	}

	tags = strings.Split(splittable, " ")
	return
}

/*
 * Get the number of tags of a type of this post
 *
 * artist: 		tags of the artists who made this post
 * character: 	tags of the characters in this post
 * copyright: 	tags of the source material for this post
 * general: 	tags that describe the content of this post
 * meta: 		tags that describe the properties of this post
 * all: 		all tags of all types
 */
func (post Post) TagsCount(tag_type string) (count uint) {
	switch tag_type {
	case "artist":
		count = post.tag_count_artist
		return
	case "character":
		count = post.tag_count_character
		return
	case "copyright":
		count = post.tag_count_copyright
		return
	case "general":
		count = post.tag_count_character
		return
	case "meta":
		count = post.tag_count_meta
		return
	case "all":
	default:
		count = post.tag_count
		return
	}
	return
}
