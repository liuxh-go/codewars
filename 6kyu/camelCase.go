package _kyu

import "strings"

/*
	将输入字符串按空格拼接成大驼峰字符串
*/

func CamelCase(s string) string {
	return strings.Replace(strings.Title(s), " ", "", -1)
}
