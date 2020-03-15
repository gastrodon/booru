package types

import (
	"github.com/gastrodon/booru/util"

	"time"
)

/*
 Represents a comment on a post on danbooru
*/
type Comment struct {
	Client    Client
	ID        int `json:"id"`
	PostID    int `json:"post_id"`    // The id of the post that this comment is on
	CreatorID int `json:"creator_id"` // Comment updater id
	UpdaterID int `json:"updater_id"` // Comment uploader id

	Body string `json:"body"` // Text content of this comment

	DoNotBumpPost bool `json:"do_not_bump_post"` // Did this comment bump the post that it is on?
	IsDeleted     bool `json:"is_deleted"`       // Was this comment deleted?
	IsSticky      bool `json:"is_sticky"`        // Was this comment stickied?

	CreatedDateString *string `json:"created_at"` // Formatted string of the created datetime
	UpdatedDateString *string `json:"updated_at"` // Formatted string of the last update datetime
}

func (comment Comment) CreatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(comment.CreatedDateString)
	return
}

func (comment Comment) UpdatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(comment.UpdatedDateString)
	return
}

func (comment Comment) GetPost() (post Post, exists bool, err error) {
	post, exists, err = comment.Client.GetPost(comment.PostID)
	return
}

func (comment Comment) GetCreator() (user User, exists bool, err error) {
	user, exists, err = comment.Client.GetUser(comment.CreatorID)
	return
}

func (comment Comment) GetUpdater() (user User, exists bool, err error) {
	user, exists, err = comment.Client.GetUser(comment.UpdaterID)
	return
}
