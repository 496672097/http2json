package http2json

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

//作者：limanman233
//时间 2023/12/8 19:20
//作用 ： http请求的封装

// Http2Json http请求的结构体
type Http2Json struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    any
	Proxy   string
	Client  *http.Client
}

// 设置默认值
func (h *Http2Json) setDefaultInfo() {
	if h.Method == "" {
		h.Method = "GET"
	}
	if h.Headers == nil {
		h.Headers = make(map[string]string)
		h.Headers["Content-Type"] = "application/json"
		h.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0"
	}
	if h.Client == nil {
		h.Client = &http.Client{}
	}
}

// httpRequest creates an HTTP request and returns the response headers, body, and any error encountered.
// @param method: The HTTP method to use (GET, POST, PUT, DELETE, etc.)
// @param headers: The HTTP headers to send with the request
// @param bodyData: The body data to send with the request
// @return respHeaders: The response headers
// @return respBody: The response body
// @return err: Any error encountered
func (h *Http2Json) HttpRequest(opts ...Option) (respHeaders map[string]string, respBody []byte, issucces bool, err error) {
	h.setDefaultInfo() // 设置默认值
	issucces = false
	// Convert the body data to JSON
	var body io.Reader
	if h.Body != nil {
		jsonData, err := json.Marshal(h.Body)
		if err != nil {
			return nil, nil, false, err
		}
		body = bytes.NewBuffer(jsonData)
	}
	// Create the HTTP request
	req, err := http.NewRequest(h.Method, h.Url, body)
	if err != nil {
		return nil, nil, false, err
	}
	// Set the request headers
	for key, value := range h.Headers {
		req.Header.Set(key, value)
	}
	// Perform the HTTP request

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, nil, false, err
	}
	defer resp.Body.Close()

	//获取状态码
	if resp.StatusCode == 200 {
		issucces = true
	}

	// Read the response body
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, false, err
	}

	// Convert the response headers to map[string]string
	respHeaders = make(map[string]string)
	for key, values := range resp.Header {
		respHeaders[key] = values[0]
	}

	// Return the headers, body, and no error
	return respHeaders, respBody, issucces, nil
}
