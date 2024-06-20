package http2json

import "fmt"

func (h *Http2Json) DebugPrint() {
	fmt.Println("请求的方法是：", h.Method)
	fmt.Println("请求的URL是：", h.Url)
	fmt.Println("请求的头部是：", h.Headers)
	fmt.Println("请求的Body是：", h.Body)
	fmt.Println("请求的Client是：", h.Client)
}
