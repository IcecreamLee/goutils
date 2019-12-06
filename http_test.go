package goutils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

func TestHttpGet(t *testing.T) {
	resp := HttpGet("http://api.map.baidu.com/location/ip")
	if resp == "" {
		t.Error("no response")
	}
	fmt.Println(resp)
}

func TestHttpPost(t *testing.T) {
	postMap := map[string]interface{}{"a": "123"}
	postBody, _ := json.Marshal(postMap)
	resp := HttpPost("https://dev.yichefu.cn", string(postBody), HttpContentTypeJson)
	if resp == "" {
		t.Error("no response")
	}
	fmt.Println(resp)

	postParams := url.Values{"a": {"123"}}
	resp = HttpPost("https://dev.yichefu.cn", postParams.Encode(), HttpContentTypeForm)
	if resp == "" {
		t.Error("no response")
	}
	fmt.Println(resp)
}
