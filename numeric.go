package goutils

import "strconv"

// IntToString 将传入的int64类型的值转为string类型后返回
func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
