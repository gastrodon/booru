package types

import (
	"github.com/gastrodon/booru/util"

	"strings"
	"time"
)

/*
 * Represents a media post on *booru
 */
type Post struct {
	Client Client
	ID     int `json:"id"`

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

	Source        string `json:"source"`     // Original source of this image
	PixivID       *int   `json:"pixiv_id"`   // ID of this image on pixiv
	Rating        string `json:"rating"`     // Post rating as [s]afe, [q]uestionable, or [e]xplicit
	MD5           string `json:"md5"`        // Hash of this post's media
	Score         int    `json:"score"`      // Total vote score
	UpScore       int    `json:"up_score"`   // Number of users who voted this up
	DownScore     int    `json:"down_score"` // Number of users who voted this down
	FavoriteCount int    `json:"fav_count"`  // Number of users that have favorited this post

	FileURL        string `json:"file_url"`         // URL to this post's media
	LargeFileURL   string `json:"large_file_url"`   // URL to the full file of this post's media
	PreviewFileURL string `json:"preview_file_url"` // URL to the preview file of this post's media
	FileSize       int    `json:"file_size"`        // Size of the file of this post's media
	FileType       string `json:"file_ext"`         // Type of this post's media file
	HasLarge       bool   `json:"has_large"`        // Does a large version of this post's media file exist?
	Width          int    `json:"image_width"`      // Pixel width of this post's media
	Height         int    `json:"image_height"`     // Pixel height of this post's media
	BitFlags       int    `json:"bit_flags"`        // I don't know

	UploaderID   int    `json:"uploader_id"`   // ID of the uploading user
	UploaderName string `json:"uploader_name"` // Name of the uploading user
	ApproverID   *int   `json:"approver_id"`   // ID of the approving user, if approved

	ParentID           *int    `json:"parent_id"`            // ID of the parent of this post, if any
	PoolString         *string `json:"pool_string"`          // Space separated list of pools that this post is in
	ChildrenIDsString  *string `json:"children_ids"`         // Space separated list if child ids, if any children
	HasChildren        bool    `json:"has_children"`         // Does this post have any children?
	HasVisibleChildren bool    `json:"has_visible_children"` // Does this post have any visible children?
	HasActiveChildren  bool    `json:"has_active_children"`  // Does this post have any children who's status is active?

	CreatedDateString         *string `json:"created_at"`             // Formatted string of the creation datetime
	UpdatedDateString         *string `json:"updated_at"`             // Formatted string of the last update datetime
	LastCommentDateString     *string `json:"last_commented_at"`      // Formatted string of the last comment datetime
	LastCommentBumpDateString *string `json:"last_comment_bumped_at"` // Formatted string of the last comment bump datetime
	LastNoteDateString        *string `json:"last_noted_at"`          // Formatted string of the last note datetime

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
func (post Post) CreatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(post.CreatedDateString)
	return
}

/*
* Get a time object representing some posts last update time, if any
 */
func (post Post) UpdatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(post.UpdatedDateString)
	return
}

/*
* Get a time object representing some posts last comment time, if any
 */
func (post Post) LastCommentAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(post.LastCommentDateString)
	return
}

/*
* Get a time object representing some posts last comment bump time, if any
 */
func (post Post) LastCommentBumpedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(post.LastCommentBumpDateString)
	return
}

/*
* Get a time object representing some posts last note add time, if any
 */
func (post Post) LastNoteAt() (parsed *time.Time, err error) {
	var _time time.Time
	if post.LastNoteDateString != nil {
		_time, err = time.Parse(time.RFC3339, *post.LastNoteDateString)
		parsed = &_time
	}
	return
}

func (post Post) Uploader() (uploader User, exists bool, err error) {
	uploader, exists, err = post.Client.GetUser(post.UploaderID)
	return
}

func (post Post) Approver() (approver User, exists bool, err error) {
	if post.ApproverID == nil {
		exists = false
		return
	}

	approver, exists, err = post.Client.GetUser(*post.ApproverID)
	return
}
