package utils

import (
	"time"
)

/*该方法用于将时间字符串计算多少年的时间戳*/
func AddAllCalc(s string) int64 {
	t := time.Now().Unix()
	i := String2TimeUnix(s)
	return (t-i) / 31536000
}
