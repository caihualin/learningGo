package learn_url

import (
	"testing"
	"fmt"
)

func Test_urlConvert(t *testing.T) {
	param := &queryParam{"aaaa", "AAAA", []string{"CCCC"}}
	paramMap := structToMap(param)
	fmt.Printf("原参数: %v\n", param)
	fmt.Printf("转换后: %s", urlConvert(paramMap))
}
