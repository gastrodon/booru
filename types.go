package booru

import (
	"strings"
)

/*
 * Represents a media post on *booru
 */
type Post struct {
	Client Client
	// unparsed tags
	tag_string           string `json:"tag_string"`
	tag_string_artist    string `json:"tag_string_artist"`
	tag_string_character string `json:"tag_string_character"`
	tag_string_copyright string `json:"tag_string_copyright"`
	tag_string_general   string `json:"tag_string_general"`
	tag_string_meta      string `json:"tag_string_meta"`
}

/*
 * A list of all tags on this post
 */
func (post Post) Tags() (tags []string) {
	return stirngs.Split(post.tag_string, " ")
}

/*
 * A list of artist tags. These are the artists who drew this image
 */
func (post Post) ArtistTags() (tags []string) {
	return strings.Split(post.tag_string_artist, " ")
}

/*
 * A list of character tags. These are the characters who are in this post
 */
func (post Post) CharacterTags() (tags []string) {
	return strings.Split(post.tag_string_character, " ")
}

/*
 A list of copyright tags. These describe from what ip this post is based
*/
func (post Post) CopyrightTags() (tags []string) {
	return strings.Split(post.tag_string_copyright, " ")
}

/*
 * A list of general tags. These usually describe the post's content
 */
func (post Post) GeneralTags() (tags []string) {
	return strings.Split(post.tag_string_general, " ")
}

/*
 * A list of meta tags. This includes tags such as highres and duplicate
 */
func (post Post) MetaTags() (tags []string) {
	return strings.Split(post.tag_string_meta, " ")
}
