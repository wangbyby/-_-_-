package utils

import (
	"strconv"
)

//数字转换为string

func NumToString(n int64) string {
	return strconv.FormatInt(n, 10)
}
