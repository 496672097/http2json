package http2json

import (
	"fmt"
	"reflect"
)

func (h *Http2Json) DebugPrint() {
	v := reflect.ValueOf(*h)
	t := reflect.TypeOf(*h)

	fmt.Println("Http2Json 结构体字段值:")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("字段-%s-的值为: %v\n", field.Name, value.Interface())
	}
}
