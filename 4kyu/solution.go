package _kyu

import (
	"fmt"
	"strings"
)

/*
	将输入的整数数组格式化为升序的字符串
*/

func Solution(list []int) string {
	resultList := make([]string, 0)

	for i := 0; i < len(list); {
		count, j := 1, i+1
		for ; j < len(list); j++ {
			if list[j] == list[j-1]+1 {
				count++
			} else {
				break
			}
		}

		if count < 3 {
			for index := i; index < j; index++ {
				resultList = append(resultList, fmt.Sprintf("%d", list[index]))
			}
		} else {
			resultList = append(resultList, fmt.Sprintf("%d-%d", list[i], list[j-1]))
		}

		i = j
	}

	return strings.Join(resultList, ",")
}
