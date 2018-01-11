package learn_url
// struct => url.Values => querystring

import (
	"net/url"
	"reflect"
)

type queryParam struct {
	a string
	B string
	C []string
}

// 结构体转url.Values
func structToMap(p *queryParam) url.Values {
	k := reflect.TypeOf(*p)
	v := reflect.ValueOf(*p)
	querymap := make(url.Values)
	for i := 0; i < k.NumField(); i++ {
		t := v.Field(i).Kind()
		switch t {
		case reflect.String:
			querymap[k.Field(i).Name] = append(querymap[k.Field(i).Name], v.Field(i).String())
		case reflect.Slice:
			querymap[k.Field(i).Name] = v.Field(i).Interface().([]string)
		}
	}
	return querymap
}

// url.Value转换querystring
func urlConvert(v url.Values) string{
	return v.Encode()
}
