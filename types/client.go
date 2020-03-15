package types

import (
	"github.com/gastrodon/booru/util"

	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	login       string      // username for api calls
	key         string      // key for api calls
	http_client http.Client // http client for making http requests

	Host string // base url for all api calls
}

var (
	categoryIDs map[string]int = map[string]int{
		"general":   0,
		"artist":    1,
		"copyright": 3,
		"character": 4,
	}
)

/*
 * Make a request to an endpoint of this clients host
 * query strings passed in do not need to include `login` and `key`,
 * as they are added when the map is parsed
 */
func (client Client) make_request(method, endpoint string, data io.Reader, query_strings ...map[string]string) (response *http.Response, err error) {
	var parsed_qs = util.FormatQS(query_strings...)
	var full_url string = fmt.Sprintf("%s%s.json?%s", client.Host, endpoint, parsed_qs)

	var request *http.Request
	request, err = http.NewRequest(method, full_url, data)
	if err != nil {
		return
	}

	if client.login != "" && client.key != "" {
		request.SetBasicAuth(client.login, client.key)
	}

	response, err = client.http_client.Do(request)
	return
}

/*
 * make a GET request and only return its body
 */
func (client Client) get_request_body(endpoint string, query_strings ...map[string]string) (json_bytes []byte, code int, err error) {
	var response *http.Response
	response, err = client.make_request("GET", endpoint, nil, query_strings...)
	if err == nil {
		code = response.StatusCode
		json_bytes, err = ioutil.ReadAll(response.Body)
	}

	return
}

/*
GET an api resource if it exists.
If it does not, return an empty response
*/
func (client Client) get_if_exists(endpoint string, query_strings ...map[string]string) (json_bytes []byte, exists bool, err error) {
	var code int
	json_bytes, code, err = client.get_request_body(endpoint, query_strings...)
	exists = code != 404 && code != 410
	return
}

/*
 * Give auth params to a client instance
 * This should be done before making most api calls
 */
func (client *Client) Auth(login, key string) {
	client.login = login
	client.key = key
}

/*
 * Get a post by its id
 */
func (client Client) GetPost(id int) (post Post, exists bool, err error) {
	var json_bytes []byte
	json_bytes, exists, err = client.get_if_exists(fmt.Sprintf("/posts/%d", id))
	if err == nil {
		err = json.Unmarshal(json_bytes, &post)
		post.Client = client
	}
	return
}

/*
 * Get a post that matches some md5
 */
func (client Client) GetPostMD5(md5 string) (post Post, exists bool, err error) {
	var q_strings map[string]string = map[string]string{
		"md5": md5,
	}

	var json_bytes []byte
	json_bytes, exists, err = client.get_if_exists("/posts", q_strings)
	if err == nil {
		err = json.Unmarshal(json_bytes, &post)
		post.Client = client
	}

	return
}

/*
 * Get a list of posts matching search parameters
 * tags: 	a list of tags to search for
 * raw: 	disable parsing tag alias parsing?
 * page:
 * limit:
 * random:
 */
func (client Client) GetPosts(tags []string, raw bool, page, limit int, random bool) (results []Post, err error) {
	var q_strings map[string]string = map[string]string{
		"tags": strings.Join(tags, "+"),
		"raw":  fmt.Sprintf("%t", raw),
	}

	var json_bytes []byte
	json_bytes, _, err = client.get_if_exists("/posts", q_strings, util.CommonParams(page, limit, random))
	if err != nil {
		return
	}

	err = json.Unmarshal(json_bytes, &results)

	var index int
	for index, _ = range results {
		results[index].Client = client
	}

	return
}

/*
 * Get a user by their id
 */
func (client Client) GetUser(id int) (user User, exists bool, err error) {
	var json_bytes []byte
	json_bytes, exists, err = client.get_if_exists(fmt.Sprintf("/users/%d", id))
	if err == nil {
		err = json.Unmarshal(json_bytes, &user)
		user.Client = client
	}

	return
}

/*
 * Get a list of users matching a map of search parameters
 * search: 	A map of user search terms defined by the danbooru api
 * 		   	These will be wrapped in a search[]
 * 	   	name: 			Name to search for
 * 	   	name_matches: 	Does the same thing as name
 * 	   	min_level: 		minimum level of users
 * 	   	max_level: 		maximum level of users
 * 	   	level: 			exact level of users
 * 	   	id: 			ID to search for
 * 	   	order: 			Search results order: Can be one of
 * 	   					name, post_upload_count, note_count,
 * 	   					post_update_count, date
 * page:
 * limit:
 */
func (client Client) GetUsers(search map[string]string, page, limit int) (results []User, err error) {
	search = util.WrapQS("search", search)

	var json_bytes []byte
	json_bytes, _, err = client.get_if_exists("/users", search, util.CommonParams(page, limit, false))
	if err != nil {
		return
	}

	err = json.Unmarshal(json_bytes, &results)

	var index int
	for index, _ = range results {
		results[index].Client = client
	}
	return
}

/*
 * Get a pool by its id
 */
func (client Client) GetPool(id int) (pool Pool, exists bool, err error) {
	var json_bytes []byte
	json_bytes, exists, err = client.get_if_exists(fmt.Sprintf("/pools/%d", id))
	if err == nil {
		err = json.Unmarshal(json_bytes, &pool)
		pool.Client = client
	}

	return
}

/*
 * Get the profile that this client is authed with
 */
func (client Client) GetProfile() (profile Profile, authed bool, err error) {
	var json_bytes []byte
	var code int
	json_bytes, code, err = client.get_request_body("/profile")
	authed = code == 200
	if err == nil && authed {
		err = json.Unmarshal(json_bytes, &profile)
		profile.Client = client
	}
	return
}

/*
 * Get a comment by its id
 */
func (client Client) GetComment(id int) (comment Comment, exists bool, err error) {
	var json_bytes []byte
	json_bytes, exists, err = client.get_if_exists(fmt.Sprintf("/comments/%d", id))
	if err == nil {
		err = json.Unmarshal(json_bytes, &comment)
		comment.Client = client
	}

	return
}

/*
 * Get related tags to some tag
 * tag:			the tag to query related tags of
 * category: 	can be one of any, general, artist, copyright, character
 */
func (client Client) GetRelatedTags(tag, category string) (related RelatedTagResponse, err error) {
	var query map[string]string = map[string]string{
		"query": tag,
	}

	if category != "all" {
		query["category"] = fmt.Sprintf("%d", categoryIDs[category])
	}

	var response_data []byte
	response_data, _, err = client.get_if_exists("/related_tag", query)
	if err == nil {
		err = json.Unmarshal(response_data, &related)
	}

	return
}
