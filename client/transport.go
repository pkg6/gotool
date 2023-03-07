package client

import "net/http"

func (c *Client) SetTransport(transport *http.Transport) *Client {
	c.httpClient.Transport = transport
	return c
}
