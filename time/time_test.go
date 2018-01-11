package learn_time

import (
	"fmt"
	"testing"
)

func Test_timefunc(t *testing.T) {
	fmt.Println("int64: ", stampInt64())
	fmt.Println("int64: ", stampNanoInt64())
	fmt.Println("string: ", stampString())
	fmt.Println("string: ", stampNanoString())
	fmt.Println("string: ", stampToString(stampInt64()))
	fmt.Println("int64: ", stringToStamp(stampString()))
}
