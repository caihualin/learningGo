package learn_url

import (
	"testing"
	"fmt"
)

func Test_parseUrl(t *testing.T) {
	var u urlParser
	us := urlString("http//user:password@host:port/path?b=1&a=2")
	u = &us
	u.parse()
	fmt.Printf("urlString: %s\n", u.string())
	fmt.Printf("queryString: %s\n", u.queryString())
	fmt.Printf("QueryMap: %v\n", u.queryMap())
}
