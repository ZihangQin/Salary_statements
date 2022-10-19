package utils

import (
	"fmt"
	"time"
)

/*该方法用于讲时间字符串转换为时间类型*/
func String2TimeUnix(s string) int64 {
	//获取区时
	loc,_ := time.LoadLocation("Local")
	//时间字符串转换为时间戳
	timeTime,err := time.ParseInLocation("2006-01-02",s,loc)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return timeTime.Unix()
}

