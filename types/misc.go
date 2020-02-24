package types

type RelatedTag struct {
	Value    string // Tag value
	Category int    // Tag category index of [general, artist, copyright, character]
}

type RelatedTagResponse struct {
	Category        *string      `json:"category"` // The category group of tags
	RelatedTags     []RelatedTag // Slice of related tags returned
	RelatedWikiTags []RelatedTag // Slice of related wiki tags returned

	RelatedTagsRaw     [][]interface{} `json:"tags"`
	RelatedWikiTagsRaw [][]interface{} `json:"wiki_page_tags"`
}

/*
 * Parse a related tags response so that the tags
 * are sent to RelatedTag instances
 */
func (reciever *RelatedTagResponse) ParseTags() {
	var related []RelatedTag = make([]RelatedTag, len(reciever.RelatedTagsRaw))

	var index int
	var current []interface{}
	for index, current = range reciever.RelatedTagsRaw {
		related[index] = RelatedTag{
			Value:    current[0].(string),
			Category: int(current[1].(float64)),
		}
	}

	reciever.RelatedTags = related
}

/*
 * Parse a related tags response so that the wiki tags
 * are sent to RelatedTag instances
 */
func (reciever *RelatedTagResponse) ParseWikiTags() {
	reciever.RelatedWikiTags = make([]RelatedTag, len(reciever.RelatedWikiTagsRaw))

	var index int
	var current []interface{}
	for index, current = range reciever.RelatedWikiTagsRaw {
		reciever.RelatedWikiTags[index] = RelatedTag{
			Value:    current[0].(string),
			Category: int(current[1].(float64)),
		}
	}
}
