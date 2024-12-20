package http2json

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	proxy   string
	client  *http.Client
	Errors  []error //错误信息
}

// 设置默认值
// 如果没有设置Method,Headers,client则设置默认值
func (h *Http2Json) setDefaultInfo(opts []Option) {
	if h.Method == "" {
		h.Method = "GET"
	}
	if h.Headers == nil {
		h.Headers = make(map[string]string)
		h.Headers["Content-Type"] = "application/json"
		h.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0"
	}
	if h.client == nil {
		h.client = &http.Client{}
	}
	//处理opts
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(h)
		}
	}
	if h.proxy != "" {
		proxyurl, err := url.Parse(h.proxy)
		if err != nil {
			h.Errors = append(h.Errors, err)
			return
		}
		// 设置代理并忽略证书安全问题
		h.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyurl),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
}

// httpRequest creates an HTTP request and returns the response headers, body, and any error encountered.
// @param method: The HTTP method to use (GET, POST, PUT, DELETE, etc.)
// @param headers: The HTTP headers to send with the request
// @param bodyData: The body data to send with the request
// @return respHeaders: The response headers
// @return respBody: The response body
// @return err: Any error encountered
func (h *Http2Json) HttpRequest(opts ...Option) (respHeaders map[string]string, respBody []byte, err error) {
	h.setDefaultInfo() // 设置默认值

	// 将 Body 数据转换为 JSON
	var body io.Reader
	if h.Body != nil {
		jsonData, err := json.Marshal(h.Body)
		if err != nil {
			return nil, nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest(h.Method, h.Url, body)
	if err != nil {
		return nil, nil, err
	}

	// 设置请求头
	for key, value := range h.Headers {
		req.Header.Set(key, value)
	}

	// 执行 HTTP 请求
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		// 返回已读取的部分数据和错误
		return nil, respBody, err
	}

	// 验证是否完整读取响应体
	if contentLengthStr := resp.Header.Get("Content-Length"); contentLengthStr != "" {
		contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
		if err == nil && int64(len(respBody)) != contentLength {
			return nil, respBody, fmt.Errorf("响应体读取不完整：期望 %d 字节，实际读取 %d 字节", contentLength, len(respBody))
		}
	}

	// 将响应头转换为 map[string]string
	respHeaders = make(map[string]string)
	for key, values := range resp.Header {
		respHeaders[key] = values[0]
	}

	// 返回响应头、响应体和错误信息
	return respHeaders, respBody, nil
}
