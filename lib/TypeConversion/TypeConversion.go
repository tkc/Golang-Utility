package TypeConversion

import (
		"strconv"
)

func Int64ToString(num int64) string {
		return strconv.FormatInt(num, 10)
}

func StringToInt64(str string) int64 {
		id, _ := strconv.ParseInt(str, 10, 64)
		return id
}