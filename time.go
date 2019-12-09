package goutils

import (
	"time"
)

// Datetime 获取当前时间 yyyy-MM-dd HH:mm:ss
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// dateFormat 将时间t格式化为字符串类型返回
func DateFormat(t time.Time, layout ...string) string {
	if len(layout) == 0 {
		layout = append(layout, "2006-01-02")
	}
	return t.Format(layout[0])
}

// datetimeFormat 将时间t格式化为字符串类型返回
func DatetimeFormat(t time.Time, layout ...string) string {
	if len(layout) == 0 {
		layout = append(layout, "2006-01-02 15:04:05")
	}
	return t.Format(layout[0])
}

// Timestamp 返回一个Unix时间戳
func Timestamp() int64 {
	return time.Now().Unix()
}

// UnixTime 返回一个Unix时间戳
func UnixTime() int64 {
	return time.Now().Unix()
}

// UnixMicroTime 返回一个Unix毫秒秒时间戳
func UnixMilliTime() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

// UnixMicroTime 返回一个Unix微秒时间戳
func UnixMicroTime() int64 {
	return time.Now().UnixNano() / 1000
}

// UnixMicroTime 返回一个Unix毫秒时间戳
func UnixNanoTime() int64 {
	return time.Now().UnixNano()
}
