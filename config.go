package http2json

import "net/http"

// Option 选项
// 作者：limanman233
// 时间 2024/06/20 15:32

type Option func(*Http2Json)

// WithProxy 设置代理
func WithProxy(proxy string) Option {
	return func(o *Http2Json) {
		o.proxy = proxy
	}
}

// WithHttpClient 设置httpclient
func WithHttpClient(httpClient *http.Client) Option {
	return func(o *Http2Json) {
		o.client = httpClient
	}
}

// withAuth 设置Authorization 需要自己设置bear
func WithAuth(Authorization string) Option {
	return func(o *Http2Json) {
		o.Headers["Authorization"] = Authorization
	}
}
