package types

import (
	"github.com/gastrodon/booru/util"

	"time"
)

type Pool struct {
	Client Client
	ID     int    `json:"id"`
	Name   string `json:"name"` // The name of this pool

	Category    string `json:"category"`    // The category of this pool
	Description string `json:"description"` // The full description of this pool and it's posts

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

/*
 * Get posts (start -> stop) from this pool
 * Start should be >= 0
 * Stop should <= total post count (`Pool.PostCount`)
 */
func (pool Pool) PostsRange(start, stop int) (posts []Post, err error) {
	posts = make([]Post, stop-start)

	var buf Post
	var index, id int
	for index, id = range pool.PostIDs[start:stop] {
		buf, _, err = pool.Client.GetPost(id)
		if err != nil {
			return
		}

		buf.Client = pool.Client
		posts[index] = buf
	}

	return
}

func (pool Pool) Posts() (posts []Post, err error) {
	posts, err = pool.PostsRange(0, pool.PostCount)
	return
}

func (pool Pool) PostAt(index int) (post Post, exists bool, err error) {
	post, exists, err = pool.Client.GetPost(pool.PostIDs[index])
	return
}
