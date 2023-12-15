package utils

import (
	"reflect"
)

// 将一个结构体的所有参数名和类型转换为一个slice
func StructToSlice(param any) []string {
	t := reflect.TypeOf(param)
	data := []string{}
	if t.Kind() == reflect.Struct {
		count := t.NumField()
		for i := 0; i < count; i++ {
			field := t.Field(i)
			data = append(data, field.Tag.Get("json"))
		}
	}

	return data
}

// 将一个结构体的所有参数名和类型转换为一个map
func StructToMap(param any) map[string]string {
	t := reflect.TypeOf(param)

	data := map[string]string{}

	if t.Kind() == reflect.Struct {
		count := t.NumField()
		for i := 0; i < count; i++ {
			field := t.Field(i)
			data[field.Tag.Get("json")] = field.Tag.Get("type")
		}
	}

	return data
}
