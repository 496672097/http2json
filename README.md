## 使用http2.json生成一个类。发送设置header body url即可


`示例代码：`

```golang
func main() {
	type Infor struct {
		Username     string `json:"Username"`
		Email        string `json:"Email"`
		PasswordHash string `json:"PasswordHash"`
	}
	http2json := http2json.Http2Json{}
	http2json.Body = Infor{Username: "admin", Email: "123@qq.com", PasswordHash: "123456"}
	http2json.Url = "https://api.lmm233.com/gettoken"
	http2json.Method = "POST"
	_, respBody, err := http2json.HttpRequest()
	if err != nil {
		println(err.Error())
        return
	}
    fmt.Println(string(respBody))
}
```