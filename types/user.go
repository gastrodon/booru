package types

import (
	"github.com/gastrodon/booru/util"

	"time"
)

/*
 Represents a user on danbooru
*/
type User struct {
	Client Client
	ID     int    `json:"id"`
	Name   string `json:"name"` // This user's username

	Level     int    `json:"level"`        // This user's level
	LevelName string `json:"level_string"` // The name of this user's level

	InviterID         *int    `json:"inviter_id"` // The id of this user's inviter, if any
	CreatedDateString *string `json:"created_at"` // Formatted string of the creation datetime

	PostUpdateCount    int `json:"post_update_count"`    // The number of post updates by this user
	NoteUpdateCount    int `json:"note_update_count"`    // The number of note updates by this user
	PostUploadCount    int `json:"post_upload_count"`    // The number of posts created by this user
	CommentCount       int `json:"comment_count"`        // The number of comments by this user
	ForumPostCount     int `json:"forum_post_count"`     // The number of forum posts by this user
	FavoriteGroupCount int `json:"favorite_group_count"` // The number of favorite groups by this user
	AppealCount        int `json:"appeal_count"`         // The number of appeals made by this account
	FlagCount          int `json:"flag_count`            // The number of report flags by this account

	PostUploadLimit int  `json:"post_upload_count"` // The current upload limit per day
	BaseUploadLimit int  `json:"base_upload_limit"` // The base upload limit per day
	MaxUploadLimit  int  `json:"max_upload_limit"`  // The maximum upload limit per day
	CanUploadFree   bool `json:"can_upload_free"`   // Can this user ignore upload limits?

	Banned     bool `json:"is_banned"`         // Is this user banned?
	SuperVoter bool `json:"is_super_voter"`    // Is this user a super voter?
	Approver   bool `json:"can_approve_posts"` // May this user approve posts?

	// The meaning of the properties below are unknown to me
	ArtistVersionCount           int `json:"artist_version_count"`
	ArtistCommentaryVersionCount int `json:"artist_commentary_version_count"`
	WikiPageVersionCount         int `json:"wiki_page_version_count"`
	PoolVersionCount             int `json:"pool_version_count"`
	PositiveFeedbackCount        int `json:"positive_feedback_count"`
	NegativeFeedbackCount        int `json:"negative_feedback_count"`
	NeutralFeedbackCount         int `json:"neutral_feedback_count"`
}

/*
 Get a time object representing this account's creation date
*/
func (user User) CreatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(user.CreatedDateString)
	return
}
