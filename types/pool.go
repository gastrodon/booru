package types

import (
	"github.com/gastrodon/booru/util"

	"time"
)

type Pool struct {
	Client Client
	ID     int    `json:"id"`
	Name   string `json:"name"` // The name of this pool

	CreatorID   int    `json:"creator_id"`   // ID if the pool's creator
	CreatorName string `json:"creator_name"` // Name of the pool's creator
	Category    string `json:"category"`     // The category of this pool
	Description string `json:"description"`  // The full description of this pool and it's posts

	Active  bool `json:"is_active"`  // Is this pool active?
	Deleted bool `json:"is_deleted"` // Is this pool deleted?

	PostCount int   `json:"post_count"` // The number of posts in this pool
	PostIDs   []int `json:"post_ids"`   // The IDs of the posts in this pool

	CreatedDateString *string `json:"created_at"` // Formatted string of the creation datetime
	UpdatedDateString *string `json:"updated_at"` // Formatted string of the last update datetime
}

/*
 * Get a time object representing this pool's creation date
 */
func (pool Pool) CreatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(pool.CreatedDateString)
	return
}

/*
* Get a time object representing some pools last update time, if any
 */
func (pool Pool) UpdatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(pool.UpdatedDateString)
	return
}

func (pool Pool) Posts() (posts []Post, err error) {
	posts = make([]Post, pool.PostCount)

	var buf Post
	var index, id int
	for index, id = range pool.PostIDs {
		buf, _, err = pool.Client.GetPost(id)
		if err != nil {
			return
		}

		buf.Client = pool.Client
		posts[index] = buf
	}

	return
}

func (pool Pool) PostAt(index int) (post Post, exists bool, err error) {
	post, exists, err = pool.Client.GetPost(pool.PostIDs[index])
	return
}
