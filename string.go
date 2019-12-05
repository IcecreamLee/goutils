package goutils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// ToSnakeCase将传入的驼峰字符串str转为蛇形字符串风格返回
func ToSnakeCase(str string) string {
	buffer := NewBuffer()
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// ToSnakeCase将传入的蛇形字符串str转为驼峰字符串返回
func ToCamelCase(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Title(str)
	return strings.Replace(str, " ", "", -1)
}

// ToString将传入的值转为字符串类型后返回
func ToString(str interface{}) string {
	if values, ok := str.([]interface{}); ok {
		var results []string
		for _, value := range values {
			results = append(results, ToString(value))
		}
		return strings.Join(results, "_")
	} else if bytes, ok := str.([]byte); ok {
		return string(bytes)
	} else if reflectValue := reflect.Indirect(reflect.ValueOf(str)); reflectValue.IsValid() {
		return fmt.Sprintf("%v", reflectValue.Interface())
	}
	return ""
}

// ToInt 将传入的值转为int类型后返回，如转换失败则返回0
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = 0
	}
	return i
}

// ToInt64 将传入的值转为int64类型后返回，如转换失败则返回0
func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = 0
	}
	return i
}
