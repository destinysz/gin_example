package utils

import (
	"reflect"
)

func StructToMapViaReflect(s interface{}) map[string]interface{} {
	// 通过反射的形式转换成map，是通过转成json，再json转成map的两倍。并且不会过滤掉零值
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
