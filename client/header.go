package client

import (
	"encoding/base64"
	"github.com/pkg6/gotool/types"
)

// WithBasicAuth 携带Authorization
func (c *Client) WithBasicAuth(username, password string) *Client {
	if c.debug {
		c.log.Debug("with BasicAuth username: %s ,password:", username, password)
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
		c.log.Debug("set Token Key=%s value", tokenType, token)
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
	c.Header.Set(key, value)
	return c
}

// WithUserAgent 携带User-Agent
func (c *Client) WithUserAgent(userAgent string) *Client {
	c.Header.Set(headerUserAgentKey, userAgent)
	return c
}

// 如果设置Content-Type就不需要进行覆盖
func (c *Client) header(key, value string) *Client {
	c.Header.SetForce(key, value, false)
	return c
}
