package learn_url
// url path的编码解码

import (
	"net/url"
	"strings"
)

// 标准编码格式application/x-www-form-urlencoded
func urlEncode(urlstring string) string {
	return url.QueryEscape(urlstring)
}

// 解码urlencoded
func urlDecode(urlEncoded string) string {
	if urlDecoded, err := url.QueryUnescape(urlEncoded); err == nil {
		return urlDecoded
	}
	return ""
}

// 替换特殊字符
func urlPercentParse(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}
