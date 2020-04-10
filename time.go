package goutils

import (
	"strconv"
	"time"
)

// Datetime 获取当前时间 yyyy-MM-dd HH:mm:ss
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// IntDate 返回当前时间的int类型的格式化日期
func IntDate() int {
	return ToInt(time.Now().Format("20060102"))
}

// IntDatetime 返回当前时间的int类型的格式化日期时间
func IntDatetime() int {
	return ToInt(time.Now().Format("20060102150405"))
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

// Timestamp 返回当前Unix时间戳
func Timestamp() int64 {
	return time.Now().Unix()
}

// UnixTime 返回当前Unix时间戳
func UnixTime() int64 {
	return time.Now().Unix()
}

// UnixMicroTime 返回当前Unix毫秒秒时间戳
func UnixMilliTime() int64 {
	return time.Now().UnixNano() / 1000000
}

// UnixMicroTime 返回当前Unix微秒时间戳
func UnixMicroTime() int64 {
	return time.Now().UnixNano() / 1000
}

// UnixMicroTime 返回当前Unix纳秒时间戳
func UnixNanoTime() int64 {
	return time.Now().UnixNano()
}

// Int2datetime 返回一个格式化的日期字符串,
// intDate 为需要格式化的int类型日期值,例如20200101,
// layout 为日期格式化的格式,例如2006-01-02.
func Int2date(intDate int, layout ...string) string {
	intDateStr := strconv.Itoa(intDate)
	t := Datetime2time(intDateStr)
	return DateFormat(t, layout...)
}

// Int2datetime 返回一个格式化的日期字符串,
// intDatetime 为需要格式化的int类型日期值,例如20200101000000,
// layout 为日期格式化的格式,例如2006-01-02 15:04:05.
func Int2datetime(intDatetime int, layout ...string) string {
	intDateStr := strconv.Itoa(intDatetime)
	t := Datetime2time(intDateStr)
	return DatetimeFormat(t, layout...)
}

// Datetime2time 返回一个time.Time类型的值,
// 如果datetime有值，则使用datetime的值转换为time.Time类型,
// 如果为传入datetime值，则使用当前时间的time.Time.
func Datetime2time(datetime ...string) time.Time {
	if len(datetime) == 0 {
		return time.Now()
	}
	dt := datetime[0]
	layout := SubStr("2006-01-02 15:04:05", 0, len(dt))
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(layout, datetime[0], loc)
	if err == nil {
		return t
	}
	return time.Now()
}
