package types

import (
	"github.com/gastrodon/booru/util"

	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
func (client Client) get_request_body(endpoint string, query_strings map[string]string) (json_bytes []byte, err error) {
	var response *http.Response
	response, err = client.make_request("GET", endpoint, query_strings, nil)
	if err == nil {
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
func (client Client) GetPost(id int) (post Post, err error) {
	var response_data []byte
	response_data, err = client.get_request_body(fmt.Sprintf("/posts/%d", id), map[string]string{})

	if err == nil {
		err = json.Unmarshal(response_data, &post)
	}

	return
}

/*
 * Get a user by their id
 */
func (client Client) GetUser(id int) (user User, err error) {
	var response_data []byte
	response_data, err = client.get_request_body(fmt.Sprintf("/users/%d", id), map[string]string{})

	if err == nil {
		err = json.Unmarshal(response_data, &user)
	}

	return
}
