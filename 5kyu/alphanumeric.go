package _kyu

import (
	"regexp"
)

/*
	验证输入的字符串是否只有字母和数字构成
*/

func Alphanumeric(str string) bool {
	isMatch, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)

	return isMatch
}
