package client

import (
	"github.com/pkg6/gotool/logger"
	"github.com/pkg6/gotool/types"
	"io"
	"net/http"
	"strings"
	"time"
)

func (c *Client) SetClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	c.httpClient = client
	return c
}

func (c *Client) SetLog(w io.Writer) *Client {
	logger.SetOutput(w)
	return c
}
func (c *Client) SetBaseURL(url string) *Client {
	c.baseURL = strings.TrimRight(url, "/")
	return c
}
func (c *Client) SetTimeOut(timeout int) *Client {
	c.httpClient.Timeout = time.Duration(timeout) * time.Second
	return c
}

// SetQuerys 设置url请求参数
func (c *Client) SetQuerys(params types.MapStrings) *Client {
	for p, v := range params {
		c.SetQuery(p, v)
	}
	return c
}

// SetQuery 设置url请求参数
func (c *Client) SetQuery(key, value string) *Client {
	c.query.Set(key, value)
	return c
}
