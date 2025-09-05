package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type curl struct {
	data        map[string]string
	method      string
	uri         string
	baseUrl     string
	header      map[string]string
	cookie      map[string]string
	contentType string
}

func Curl(baseUrl string) *curl {
	return &curl{baseUrl: baseUrl}
}

func (c *curl) SetCookie(cookie map[string]string) *curl {
	c.cookie = cookie
	return c
}

func (c *curl) SetHeader(header map[string]string) *curl {
	c.header = header
	return c
}

func (c *curl) SetData(data map[string]string) *curl {
	c.data = data
	return c
}

// SetType 类型，json还是form
// Type = "json" or "form" , 默认form
func (c *curl) SetType(Type string) *curl {
	c.contentType = Type
	return c
}

func (c *curl) Get(uri string) string {

	resp, _ := http.Get(c.getGetUrl(uri))
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

func (c *curl) Run(uri, method string) (string, error) {
	client := &http.Client{}
	targetUrl := c.getUrl(uri)

	data := c.getData()
	if strings.ToUpper(method) == "GET" {
		data = nil
		targetUrl = c.getGetUrl(uri)
	}

	req, err := http.NewRequest("POST", targetUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	Type := "application/x-www-form-urlencoded"
	if strings.ToLower(c.contentType) == "json" {
		Type = "application/json"
	}

	//设置请求头
	req.Header.Set("content-type", Type)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")

	if len(c.header) > 0 {
		for k, v := range c.header {
			req.Header.Add(k, v)
		}
	}

	if len(c.cookie) > 0 {
		for k, v := range c.header {
			cookie := &http.Cookie{Name: k, Value: v}
			req.AddCookie(cookie)
		}
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, _ := io.ReadAll(resp.Body)

	return string(body), nil
}

func (c *curl) getData() url.Values {
	data := url.Values{}
	if len(c.data) > 0 {
		for k, v := range c.data {
			data.Add(k, v)
		}
	}

	return data
}

func (c *curl) getUrl(uri string) string {
	return c.baseUrl + uri
}

func (c *curl) getGetUrl(uri string) string {
	targetUrl := c.getUrl(uri)

	if len(c.data) > 0 {
		u, _ := url.ParseRequestURI(targetUrl)

		// URL param
		data := c.getData()

		u.RawQuery = data.Encode() // URL encode

		return u.String()
	}

	return targetUrl
}
