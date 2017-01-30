package TypeConversion

import (
		"strconv"
)

func IntToString(num int) string {
		return strconv.Itoa(num)
}

func Int64ToString(num int64) string {
		return strconv.FormatInt(num, 10)
}

func StringToInt(str string) int {
		number, _ := strconv.Atoi(str)
		return number
}

func StringToInt64(str string) int64 {
		id, _ := strconv.ParseInt(str, 10, 64)
		return id
}

func IntToInt64(num int) int64 {
		i64 := int64(num)
		return i64
}
