package types

import (
	"github.com/gastrodon/booru/util"

	"time"
)

/*
 Represents the authed client on danbooru
*/
type Profile struct {
	Client             Client
	Level              int    `json:"level"`                // This profile's level
	LevelString        string `json:"level_string"`         // This profile's level's string
	ID                 int    `json:"id"`                   // The id that points to this profile's user
	InviterID          int    `json:"inviter_id"`           // The id of the user who invited this profile, can also be self
	ApiRegenMultiplier int    `json:"api_regen_multiplier"` // API quota regeneraton multiplier

	PerPage            int     `json:"per_page"`             // Default number of items to get per page in paginated requests
	MaxSavedSearches   int     `json:"max_saved_searches"`   // Save searches cap
	MaxUploadLimit     int     `json:"max_upload_limit"`     // Upload limit cap
	UploadLimit        int     `json:"upload_limit"`         // Current upload limit
	ApiBurstLimit      int     `json:"api_burst_limit"`      // API burst usage limit
	BaseUploadLimit    int     `json:"base_upload_limit"`    // Base max uploads
	TagQueryLimit      int     `json:"tag_query_limit"`      // Max tags per search
	FavoriteGroupLimit int     `json:"favorite_group_limit"` // Max favorite groups
	FavoriteLimit      int     `json:"favorite_limit`        // Max favorite posts
	StatementTimeout   int     `json:"statement_timeout"`
	CommentThreshold   int     `json:"comment_threshold"`   // Minimum comment score to be displayed by default
	RemainingApiLimit  float64 `json:"remaining_api_limit"` // API quota remaining. Why is it a float? I don't know

	FlagCount             int `json:"flag_count"`       // Number of posts flagged
	ForumPostCount        int `json:"forum_post_count"` // Number of forum posts made
	PositiveFeedbackCount int `json:"positive_feedback_count"`
	NegativeFeedbackCount int `json:"negative_feedback_count"`
	NeutralFeedbackCount  int `json:"neutral_feedback_coun"`
	PostUpdateCount       int `json:"post_update_count"`    // Number of posts updated
	NoteUpdateCount       int `json:"note_update_count"`    // Number of post notes updated
	PostUploadCount       int `json:"post_upload_count"`    // Number of posts uploaded
	FavoriteCount         int `json:"favorite_count"`       // Number of posts favorited`
	FavoriteGroupCount    int `json:"favorite_group_count"` // Number of favorite groups created`
	CommentCount          int `json:"comment_count"`        // Number of comments left on posts
	AppealCount           int `json:"appeal_count"`         // Number of appeals made

	WikiPageVersionCount         int `json:"wiki_page_version_count"`
	PoolVersionCount             int `json:"pool_version_count"`
	ArtistVersionCount           int `json:"artist_version_count"`
	ArtistCommentaryVersionCount int `json:"artist_commentary_version_count"`

	CustomStyle           string `json:"custom_style"`
	Email                 string `json:"email"`              // Email address attached to this profile
	Theme                 string `json:"theme"`              // Default website theme
	Name                  string `json:"name"`               // Then name of this profile
	DefaultImageSize      string `json:"default_image_size"` // Default image display size
	BlacklistedTagsString string `json:"blacklisted_tags"`   // Tags blacklisted in normal search
	FavoriteTagsString    string `json:"favorite_tags"`      // Favorited tags

	TimeZoneString          string  `json:"time_zone"`          // Profile timezone
	UpdatedDateString       *string `json:"updated_at"`         // When this profile was last updated
	CreatedDateString       *string `json:"created_at"`         // When this profile was created
	LastForumReadDateString *string `json:"last_forum_read_at"` // When the profile last read a forum post
	LastLoggedInDateString  *string `json:"last_logged_in_at"`  // When this profile was last logged in

	CanApprove         bool `json:"can_approve_posts"`     // May this profile approve posts?
	CanComment         bool `json:"can_comment"`           // May this profile create comments?
	CanCommentVote     bool `json:"can_comment_vote"`      // May this profile vote on comments?
	CanRemoveFromPools bool `json:"can_remove_from_pools"` // May this profile remove posts from pools?
	CanUpload          bool `json:"can_upload"`            // May this profile upload images?
	CanUploadFree      bool `json:"can_upload_free"`       // May this profile upload images without api ratelimits?
	IsBanned           bool `json:"is_banned"`             // Is this profile banned?
	IsCommentLimited   bool `json:"is_comment_limited"`    // Is this profile's commenting limited?
	IsSuperVoter       bool `json:"is_super_voter"`        // Is this profile a super voter?

	DisableCategorizedSavedSearches bool `json:"disable_categorized_saved_searches"`
	DisableCroppedThumbnails        bool `json:"disable_cropped_thumbnails"`
	DisableMobileGestures           bool `json:"disable_mobile_gestures"`
	DisablePostTooltips             bool `json:"disable_post_tooltips"`
	DisableResponsiveMode           bool `json:"disable_responsive_mode"`
	DisableTaggedFilenames          bool `json:"disable_tagged_filenames"`

	EnableAutoComplete             bool `json:"enable_auto_complete"`
	EnablePostNavigation           bool `json:"enable_post_navigation"`
	EnablePrivateFavorites         bool `json:"enable_private_favorites"`
	EnableRecentSearches           bool `json:"enable_recent_searches"`
	EnableRecommendedPosts         bool `json:"enable_recommended_posts"`
	EnableSafeMode                 bool `json:"enable_safe_mode"`
	EnableSequentialPostNavigation bool `json:"enable_sequential_post_navigation"`

	HasSavedSearches          bool `json:"has_saved_searches"`          // Does this profile have saved searches?
	HasMail                   bool `json:"has_mail"`                    // Does this profile have new mail?
	ReceiveEmailNotifications bool `json:"receive_email_notifications"` // Does this profile recieve email notifications?

	AlwaysResizeImages      bool `json:"always_resize_images"`       // Always resize images when viewing?
	HideDeletedPosts        bool `json:"hide_deleted_posts"`         // Hide deleted posts when searching?
	NewPostNavigationLayout bool `json:"new_post_navigation_layout"` // Use new post navigation layout?
	OptOutTracking          bool `json:"opt_out_tracking"`           // Has this profile opted out of tracking?
	ShowDeletedChildren     bool `json:"show_deleted_children"`      // Show deleted children when viewing posts?
	StyleUsernames          bool `json:"style_usernames"`            // Apply styling to usernames?
	NoFeedback              bool `json:"no_feedback"`
	NoFlagging              bool `json:"no_flagging"`
}

func (profile Profile) UpdatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(profile.UpdatedDateString)
	return
}

func (profile Profile) CreatedAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(profile.CreatedDateString)
	return
}

func (profile Profile) LastForumReadAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(profile.LastForumReadDateString)
	return
}

func (profile Profile) LastLoggedInAt() (parsed *time.Time, err error) {
	parsed, err = util.TimeFromPtr(profile.LastLoggedInDateString)
	return
}

/*
 * Get the user attached to this profile
 */
func (profile Profile) GetUser() (user User, exists bool, err error) {
	user, exists, err = profile.Client.GetUser(profile.ID)
	return
}
