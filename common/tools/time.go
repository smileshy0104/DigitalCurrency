package tools

import "time"

// ISO 将时间转换为RFC3339格式的字符串。
// 该函数接收一个time.Time类型的参数t，将其转换为UTC时间，并按照RFC3339格式进行格式化。
func ISO(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// ToMill 将字符串形式的时间转换为毫秒级时间戳。
// 该函数接收一个字符串参数str，其格式应为"2006-01-02 15:04:05"。
// 函数将该字符串解析为时间，并返回自Unix纪元（1970-01-01 00:00:00 UTC）以来的毫秒数。
func ToMill(str string) int64 {
	parse, _ := time.Parse("2006-01-02 15:04:05", str)
	return parse.UnixMilli()
}

// ToTimeString 将毫秒级时间戳转换为字符串形式的时间。
// 该函数接收一个int64类型的参数mill，代表自Unix纪元以来的毫秒数。
// 函数将其转换为时间，并以"2006-01-02 15:04:05"格式返回字符串表示的时间。
func ToTimeString(mill int64) string {
	milli := time.UnixMilli(mill)
	return milli.Format("2006-01-02 15:04:05")
}

// ZeroTime 返回当天零点的时间戳。
// 该函数获取当前时间，并创建一个当天零点的时间对象。
// 然后，函数返回该时间自Unix纪元以来的毫秒数。
func ZeroTime() int64 {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return date.UnixMilli()
}
