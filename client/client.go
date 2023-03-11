package client

import (
	"bytes"
	"encoding/json"
	"github.com/pkg6/gotool/logger"
	"github.com/pkg6/gotool/types"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func New(baseURL string, fns ...func(client *Client)) *Client {
	client := Client{}.Clone()
	client.SetBaseURL(baseURL)
	if client.httpClient == nil {
		client.SetClient(nil)
	}
	for _, fn := range fns {
		fn(client)
	}
	return client
}

type Client struct {
	debug bool
	// base url
	baseURL string
	// query
	query types.MapStrings
	//header
	headers types.MapStrings
	//cookie
	cookies types.MapStrings
	//必须需要初始化
	httpClient *http.Client
	//最终执行的生成的url
	URL *url.URL
	//响应
	Response *http.Response
}

func (c Client) Clone() *Client {
	c.debug = false
	c.baseURL = ""
	c.query = types.MapStrings{}
	c.headers = types.MapStrings{}
	c.cookies = types.MapStrings{}
	c.httpClient = nil
	c.URL, _ = url.Parse("")
	c.Response = nil
	logger.SetPrefix("GoTool Client ")
	return &c
}
func (c *Client) Debug() *Client {
	c.debug = true
	return c
}

// BuildUrl 生成完成的url参数复制给URL
func (c *Client) buildUrl(maps ...types.MapStrings) *url.URL {
	Url, _ := url.Parse(c.baseURL)
	q := Url.Query()
	for _, m := range maps {
		for k, v := range m {
			q.Set(k, v)
		}
	}
	Url.RawQuery = q.Encode()
	c.URL = Url
	return c.URL
}

// Get get请求
func (c *Client) Get(query types.MapStrings) ([]byte, error) {
	c.buildUrl(c.query, query)
	return c.Do(http.MethodGet, c.URL.String(), nil, nil, nil)
}

// FileInfo 上传文件基本信息
type FileInfo struct {
	//提交的时候，服务端需要对应一个字段名称
	Name string
	// 读取文件 os.Open("text.go")
	File *os.File
}

// PostFiles 多文件上传
func (c *Client) PostFiles(files []FileInfo, params types.MapStrings) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for _, file := range files {
		stat, err := file.File.Stat()
		if err != nil {
			return nil, err
		}
		if c.debug {
			logger.Debug("Client.Upload FileInfo", stat)
		}
		fileWriter, err := bodyWriter.CreateFormFile(file.Name, stat.Name())
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fileWriter, file.File)
		if err != nil {
			return nil, err
		}
	}
	c.SetHeader(HeaderContentTypeKey, bodyWriter.FormDataContentType())
	_ = bodyWriter.Close()
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}
	c.buildUrl(c.query)
	return c.Do(http.MethodPost, c.URL.String(), bodyBuf, nil, nil)
}

// PostFile 单文件上传
// file, _ := os.OpenFile("test.md", os.O_RDONLY, os.ModePerm)
// defer file.Close()
// UploadOne("file",file,nil)
func (c *Client) PostFile(name string, file *os.File, params types.MapStrings) ([]byte, error) {
	return c.PostFiles([]FileInfo{{Name: name, File: file}}, params)
}

// PostForm 表单提交
func (c *Client) PostForm(params url.Values) ([]byte, error) {
	c.header(HeaderContentTypeKey, FormContentType)
	c.buildUrl(c.query)
	return c.Do(http.MethodPost, c.URL.String(), strings.NewReader(params.Encode()), nil, nil)
}

// PostJson json提交
func (c *Client) PostJson(body any) ([]byte, error) {
	c.header(HeaderContentTypeKey, JsonContentType)
	var jsonStr string
	var verify bool
	if str, ok := body.(string); ok {
		jsonStr = str
		verify = true
	} else {
		if jsonByte, err := json.Marshal(body); err == nil {
			jsonStr = string(jsonByte)
			verify = true
		}
	}
	if verify {
		c.buildUrl(c.query)
		return c.Do(http.MethodPost, c.URL.String(), strings.NewReader(jsonStr), nil, nil)
	}
	return nil, nil
}
