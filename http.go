package goutils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// HttpGet 发送一个GET请求，并且返回请求响应的内容
// 如果返回为空字符串，则代表请求发生错误或者请求失败（所以不要使用此函数去请求无响应体的接口）
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

const HttpContentTypeJson string = "application/json"
const HttpContentTypeForm string = "application/x-www-form-urlencoded"

// HttpPostJson 发送一个POST请求，请求体类型为json，并且返回请求响应的内容
// contentType通常为 application/json 或者 application/x-www-form-urlencoded
// 如果返回为空字符串，则代表请求发生错误或者请求失败（所以不要使用此函数去请求无响应体的接口）
func HttpPost(postUrl string, postBody string, contentType string) string {
	resp, err := http.Post(postUrl, contentType, strings.NewReader(postBody))
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
