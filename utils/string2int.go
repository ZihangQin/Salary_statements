package utils

import "strconv"

/*该方法用于将string转化为int*/
func String2Int(s string) (int,error) {
	return strconv.Atoi(s)
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}

func Float2String(f float64) string {
	return strconv.FormatFloat(f,'f',3,64)
}
