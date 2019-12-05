package goutils

import (
	"io/ioutil"
	"net/http"
)

// HttpGet 发送一个GET请求，并且返回请求响应的内容
func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
