package goutils

import (
	"net/http"
	"strings"
)

// GetSite 返回指定请求r的域名: [协议类型]://[服务器地址]:[端口号]/, 如https://google.com/
func GetSite(r *http.Request) string {
	if r.TLS == nil {
		return "http://" + r.Host + "/"
	}
	return "https://" + r.Host + "/"
}

// GetPort 返回指定请求r的域名端口，如80
func GetPort(r *http.Request) int {
	pos := strings.Index(r.Host, ":")
	if pos > 0 {
		return ToInt(r.Host[pos+1:])
	}
	if r.TLS == nil {
		return 80
	}
	return 443
}

// GetFullURL 返回指定请求r的完整URL地址: [协议类型]://[服务器地址]:[端口号]/[路径]?[查询参数]，如：https://www.a.com/b/c?d=e
func GetFullURL(r *http.Request) string {
	return GetSite(r) + r.RequestURI[1:]
}

// GetURL 返回指定请求r不包含查询参数的URL地址: [协议类型]://[服务器地址]:[端口号]/[路径]，如：https://www.a.com/b/c
func GetURL(r *http.Request) string {
	pos := strings.Index(r.RequestURI, "?")
	if pos >= 0 {
		return GetSite(r) + r.RequestURI[1:pos]
	}
	return GetSite(r) + r.RequestURI[1:]
}

// SetURL 根据指定请求r的当前域名和指定的路径path去设置URL
// 如: 当前r的URL为https://a.com/b/c,
// SetURL(r, "/index") => https://a.com/index
// SetURL(r, "d?e=f") => https://a.com/d?e=f
func SetURL(r *http.Request, path string) string {
	if path == "" {
		return GetURL(r)
	}
	if path[0:1] == "/" {
		return GetSite(r) + path[1:]
	}
	url := GetURL(r)
	pos := strings.LastIndex(url, "/")
	return url[0:pos+1] + path
}
