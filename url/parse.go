package learn_url
// url path的解析

import (
	"net/url"
	"fmt"
)

type urlString string

type urlParser interface {
	string() string
	parse()
	queryString() string
	queryMap() url.Values
}

var urlParsed *url.URL

// 当前url
func (u *urlString) string() string {
	return fmt.Sprintf("%s", *u)
}

// 解析url
func (u *urlString) parse() {
	urlParsed, _ = url.Parse(u.string())
}

// 获取url中querystring
func (u *urlString) queryString() string {
	return urlParsed.RawQuery
}

// 将url中querystring解析为map
func (u *urlString) queryMap() url.Values {
	return urlParsed.Query()
}
