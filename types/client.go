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
func (client Client) make_request(method, endpoint string, query_strings map[string]string, data io.Reader) (response *http.Response, err error) {
	var parsed_qs = util.FormatQS(query_strings)
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
func (client Client) get_request_body(endpoint string, query_strings map[string]string) (json_bytes []byte, code int, err error) {
	var response *http.Response
	response, err = client.make_request("GET", endpoint, query_strings, nil)
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
func (client Client) GetPosts(tags []string, page, limit int, random, raw bool) (results []Post, err error) {
	var q_strings map[string]string = map[string]string{
		"limit": fmt.Sprintf("%d", limit),
		"page":  fmt.Sprintf("%d", page),
		"tags":  strings.Join(tags, " "),
	}

	if random {
		q_strings["random"] = "true"
	}

	if raw {
		q_strings["raw"] = "true"
	}

	var response_data []byte
	response_data, _, err = client.get_request_body("/posts", q_strings)
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
