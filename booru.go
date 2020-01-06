package booru

import (
	"github.com/gastrodon/booru/types"
)

/*
 * Create a new client that talks to a booru instance at `host`
 * Host should be a valid url with a schema and port (if applicable)
 */
func ClientAt(host string) (client types.Client) {
	client = types.Client{
		Host: host,
	}

	return
}
