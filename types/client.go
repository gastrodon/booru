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
	var response_data []byte
	var code int
	response_data, code, err = client.get_request_body(fmt.Sprintf("/posts/%d", id), map[string]string{})
	exists = code != 404 && code != 410
	if err == nil && exists {
		err = json.Unmarshal(response_data, &post)
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

	var response_data []byte
	var code int
	response_data, code, err = client.get_request_body("/posts", q_strings)
	exists = code != 404 && code != 410
	if err == nil && exists {
		err = json.Unmarshal(response_data, &post)
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
		"tags": strings.Join(tags, " "),
		"raw":  fmt.Sprintf("%t", raw),
	}

	var response_data []byte
	response_data, _, err = client.get_request_body("/posts", q_strings, util.CommonParams(page, limit, random))
	if err != nil {
		return
	}

	err = json.Unmarshal(response_data, &results)

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
	var response_data []byte
	var code int
	response_data, code, err = client.get_request_body(fmt.Sprintf("/users/%d", id), map[string]string{})
	exists = code != 404 && code != 410
	if err == nil && exists {
		err = json.Unmarshal(response_data, &user)
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

	var response_data []byte
	response_data, _, err = client.get_request_body("/users", search, util.CommonParams(page, limit, false))
	if err != nil {
		return
	}

	err = json.Unmarshal(response_data, &results)

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
	var response_data []byte
	var code int
	response_data, code, err = client.get_request_body(fmt.Sprintf("/pools/%d", id), map[string]string{})
	exists = code != 404 && code != 410
	if err == nil && exists {
		err = json.Unmarshal(response_data, &pool)
		pool.Client = client
	}

	return
}

/*
 * Get the profile that this client is authed with
 */
func (client Client) GetProfile() (profile Profile, authed bool, err error) {
	var response_data []byte
	var code int
	response_data, code, err = client.get_request_body("/profile", map[string]string{})
	authed = code == 200
	if err == nil && authed {
		err = json.Unmarshal(response_data, &profile)
		profile.Client = client
	}
	return
}
