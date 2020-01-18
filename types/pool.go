package types

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

	CreatedDateString string `json:"created_at"` // Formatted string of the creation datetime
	UpdatedDateString string `json:"updated_at"` // Formatted string of the last update datetime
}
