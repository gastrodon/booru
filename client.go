package booru

import (
	"github.com/gastrodon/booru/util"

	"fmt"
	"io"
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
 * Give auth params to a client instance
 * This should be done before making most api calls
 */
func (client *Client) Auth(login, key string) {
	client.login = login
	client.key = key
}

/*
 * Create a new client that talks to a booru instance at `host`
 * Host should be a valid url with a schema and port (if applicable)
 */
func ClientAt(host string) (client Client) {
	client = Client{
		Host: host,
	}

	return
}
