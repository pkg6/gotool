package client

import (
	"fmt"
	"github.com/pkg6/gotool/logger"
	"github.com/pkg6/gotool/types"
	"io"
	"io/ioutil"
	"net/http"
)

func (c *Client) HttpClient() *http.Client {
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	return c.httpClient
}

func (c *Client) HttpRequest(method, url string, body io.Reader) (*http.Request, error) {
	var err error
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return request, err
	}
	if len(c.headers) > 0 {
		for hk, hv := range c.headers {
			request.Header.Set(hk, hv)
		}
	}
	if len(c.cookies) > 0 {
		for ck, cv := range c.cookies {
			request.AddCookie(&http.Cookie{Name: ck, Value: cv})
		}
	}
	return request, nil
}

// Do 所有的请求都可以走这个方法
func (c *Client) Do(method, url string, body io.Reader, header types.MapStrings, cookie types.MapStrings) ([]byte, error) {
	var err error
	//合并header
	c.headers = types.MergeMapsString(c.headers, header)
	//合并cookie
	c.cookies = types.MergeMapsString(c.cookies, cookie)
	request, err := c.HttpRequest(method, url, body)
	if err != nil {
		logger.Error(fmt.Sprintf("c.HttpRequest err=%v", err), nil)
		return nil, err
	}
	_ = c.HttpClient()
	if c.debug {
		logger.Debug(fmt.Sprintf("Client.Do.Request %s %s", method, url), nil)
		logger.Debug("Client.Do.Request Header", c.headers)
		logger.Debug("Client.Do.Request Cookie ", c.cookies)
	}
	c.Response, err = c.httpClient.Do(request)
	if err != nil {
		logger.Error(fmt.Sprintf("client.Do.httpClient.Do err=%v", err), nil)
		return nil, err
	}
	defer c.Response.Body.Close()
	bodyByte, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Client.Request.ioutil.ReadAll err=%v", err), nil)
		return nil, err
	}
	if c.debug {
		logger.Debug("Client.Do.Response", c.Response)
		logger.Debug("Client.Do.Response Header", c.Response.Header)
		logger.Debug("Client.Do.Response Cookie ", c.Response.Cookies())
		logger.Debug("Client.Do.Response body", string(bodyByte))
	}
	return bodyByte, err
}
