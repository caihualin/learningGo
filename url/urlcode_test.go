package learn_url

import (
	"testing"
	"fmt"
)

func Test_urlEncode(t *testing.T) {
	urlString := "b=1 &a=0&~*"
	urlEncoded := urlEncode(urlString)
	urlDecoded := urlDecode(urlEncoded)
	urlPercent := urlPercentParse(urlEncoded)
	fmt.Printf("urlString: %s\n", urlString)
	fmt.Printf("urlEncode: %s\n", urlEncoded)
	fmt.Printf("urlDecode: %s\n", urlDecoded)
	fmt.Printf("urlParsePercent: %s\n", urlPercent)
}
