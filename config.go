package http2json

import "net/http"

type Option func(*Http2Json)

// WithProxy 设置代理
func WithProxy(proxy string) Option {
	return func(o *Http2Json) {
		o.Proxy = proxy
	}
}

// WithHttpClient 设置httpclient
func WithHttpClient(httpClient *http.Client) Option {
	return func(o *Http2Json) {
		o.Client = httpClient
	}
}