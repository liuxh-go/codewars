package _kyu

import (
	"fmt"
	"sort"
	"strings"
)

/*
	找到两个字符串中小写字母出现的次数
	按照出现的次数来拼接
*/

func Mix(s1, s2 string) string {
	// 统计字符串中字母出现的次数
	a, b := [26]int{}, [26]int{}
	for _, c1 := range []byte(s1) {
		if c1 >= 'a' && c1 <= 'z' {
			a[c1-'a']++
		}
	}
	for _, c2 := range []byte(s2) {
		if c2 >= 'a' && c2 <= 'z' {
			b[c2-'a']++
		}
	}

	// 组装字符串数组
	strList := make([]string, 0)
	for i := byte(0); i < 26; i++ {
		if a[i] < 2 && b[i] < 2 {
			continue
		}

		prefix, count := "", 0
		if a[i] > b[i] {
			prefix = "1"
			count = a[i]
		} else if a[i] == b[i] {
			prefix = "="
			count = a[i]
		} else {
			prefix = "2"
			count = b[i]
		}

		strList = append(strList, fmt.Sprintf("%s:%s", prefix,
			strings.Repeat(string(i+'a'), count)))
	}

	// 按长度、前缀、字母排序
	sort.Slice(strList, func(i, j int) bool {
		if len(strList[i]) < len(strList[j]) {
			return false
		}

		if len(strList[i]) == len(strList[j]) {
			return strList[i] < strList[j]
		}

		return true
	})

	return strings.Join(strList, "/")
}
