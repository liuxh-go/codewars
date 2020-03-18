package _kyu

import "strings"

/*
	找到第一个未重复的字符
*/

func FirstNonRepeating(str string) string {
	tmpStr := strings.ToLower(str)
	for index, c := range tmpStr {
		if strings.Count(tmpStr, string(c)) == 1 {
			return string(str[index])
		}
	}

	return ""
}
