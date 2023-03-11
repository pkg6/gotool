package client

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg6/gotool/logger"
	"github.com/pkg6/gotool/types"
)

var (
	HeaderUserAgentKey   = "User-Agent"
	HeaderContentTypeKey = "Content-Type"
	FormContentType      = "application/x-www-form-urlencoded;charset=utf-8"
	FormASCIIContentType = "application/x-www-form-urlencoded"
	JsonContentType      = "application/json; charset=utf-8"
	JsonpContentType     = "application/javascript; charset=utf-8"
	JsonASCIIContentType = "application/json"
)

// WithUserAgent 携带User-Agent
func (c *Client) WithUserAgent(userAgent string) *Client {
	c.SetHeader(HeaderUserAgentKey, userAgent)
	return c
}

// WithContentType 如果设置Content-Type就不需要进行覆盖
func (c *Client) WithContentType(contentType string) *Client {
	c.SetHeader(HeaderContentTypeKey, contentType)
	return c
}

// WithBasicAuth 携带Authorization
func (c *Client) WithBasicAuth(username, password string) *Client {
	if c.debug {
		logger.Debug(fmt.Sprintf("with BasicAuth username: %s ,password:%s", username, password), nil)
	}
	c.WithToken(base64.StdEncoding.EncodeToString([]byte(username+":"+password)), "", "")
	return c
}

// WithToken 携带token
func (c *Client) WithToken(token, tokenKey, tokenType string) *Client {
	if tokenType == "" {
		token = "Basic " + token
	} else {
		token = tokenType + token
	}
	if tokenKey == "" {
		tokenKey = "Authorization"
	}
	if c.debug {
		logger.Debug(fmt.Sprintf("set Token Key=%s value=%s", tokenType, token), nil)
	}
	c.SetHeader(tokenKey, token)
	return c
}

// SetHeaders 批量设置header
func (c *Client) SetHeaders(params types.MapStrings) *Client {
	for p, v := range params {
		c.SetHeader(p, v)
	}
	return c
}

// SetHeader 单独设置header
func (c *Client) SetHeader(key, value string) *Client {
	c.headers.Set(key, value)
	return c
}

// 如果设置就不需要进行覆盖
func (c *Client) header(key, value string) *Client {
	c.headers.SetForce(key, value, false)
	return c
}
