package types

import (
	"strings"
	"time"
)

/*
 * Represents a media post on *booru
 */
type Post struct {
	Client Client
	ID     int `json:"id"`
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

	Source            string `json:"source"`     // Original source of this image
	Rating            string `json:"rating"`     // Post rating as [s]afe, [q]uestionable, or [e]xplicit
	Md5               string `json:"md5"`        // Hash of this post's media
	Score             int    `json:"score"`      // Total vote score
	UpScore           int    `json:"up_score"`   // Number of users who voted this up
	DownScore         int    `json:"down_score"` // Number of users who voted this down
	FavoriteCount     int    `json:"fav_count"`  // Number of users that have favorited this post
	CreatedDateString string `json:"created_at"` // Formatted string of post creation datetime
	UpdatedDateString string `json:"updated_at"` // Formatted string of post last update datetime
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

/*
 * Get a time object representing some posts creation time
 */
func (post Post) CreatedAt() (parsed time.Time, err error) {
	parsed, err = time.Parse(time.RFC3339, post.CreatedDateString)
	return
}

/*
* Get a time object representing some posts last update time
 */
func (post Post) UpdatedAt() (parsed time.Time, err error) {
	parsed, err = time.Parse(time.RFC3339, post.UpdatedDateString)
	return
}
