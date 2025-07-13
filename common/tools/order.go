package tools

import (
	"fmt"
	"math/rand"
	"time"
)

// Unq 生成一个基于当前时间和随机数的唯一后缀字符串。
// 这个函数通常用于在基于prefix的基础之上创建唯一的标识符。
// 参数:
//
//	prefix - 前缀字符串，用于个性化生成的唯一字符串。
//
// 返回值:
//
//	一个格式化的字符串，包含prefix、当前时间和一个随机数。
func Unq(prefix string) string {
	// 获取当前时间的Unix毫秒时间戳，用于生成唯一性部分。
	milli := time.Now().UnixMilli()
	// 生成一个随机数，进一步增加唯一性。
	intn := rand.Intn(999999)
	// 将前缀、当前时间和随机数格式化为一个字符串并返回。
	return fmt.Sprintf("%s%d%d", prefix, milli, intn)
}
