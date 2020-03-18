package _kyu

import (
	"fmt"
	"regexp"
	"strconv"
)

var MultipleOf3Regex = "^(0*1(00|11)*10*)$|^0$"

func MultipleOf3RegexFunc(str string) bool {
	matched, err := regexp.MatchString(MultipleOf3Regex, str)
	if err != nil {
		fmt.Println(err)
	}

	return matched
}

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}
