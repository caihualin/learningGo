package learn_time
// 时间格式化，时间戳转换

import (
	"time"
	"strconv"
)
func stampInt64() int64 {
	return time.Now().Unix()
}

func stampNanoInt64() int64 {
	return time.Now().UnixNano() / 1000
}

func stampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func stampNanoString() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000, 10)
}

func stampToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func stringToStamp(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}