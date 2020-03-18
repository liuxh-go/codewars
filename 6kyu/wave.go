package _kyu

import (
	"strings"
)

/*
	将输入的字符串波浪形输出
*/

func Wave(words string) []string {
	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		tmpResult, index, ifValid := words[:i], i, false
		for ; index < len(words); index++ {
			if words[index] != ' ' {
				tmpResult += strings.ToUpper(string(words[index]))
				ifValid = true
				break
			} else {
				i++
				tmpResult += " "
			}
		}
		if index < len(words)-1 {
			tmpResult += words[index+1:]
		}

		if ifValid {
			result = append(result, tmpResult)
		}
	}

	return result
}
