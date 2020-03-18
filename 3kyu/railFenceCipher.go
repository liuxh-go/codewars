package _kyu

import (
	"fmt"
	"math"
)

/*
	蛇形编码解码输入的字符串

	* n>=2

	eg.

	WEAREDISCOVEREDFLEEATONCE 3

	W       E       C       R       L       T       E
	  E   R   D   S   O   E   E   F   E   A   O   C
	    A       I       V       D       E       N

	WECRLTEERDSOEEFEAOCAIVDEN
*/

func Encode(s string, n int) string {
	l := len(s)
	if l <= n {
		return s
	}

	result, maxOffset := "", (n-1)*2

	// 第一排
	index := 0
	for ; index < l; index += maxOffset {
		result += string(s[index])
	}

	// 中间的
	for i := 1; i < n-1; i++ {
		offset := (n - 1 - i) * 2
		for index := i; index < l; {
			result += string(s[index])

			index += offset
			offset = maxOffset - offset
		}
	}

	// 最后一排
	for index := n - 1; index < l; index += maxOffset {
		result += string(s[index])
	}

	return result
}

func Decode(s string, n int) string {
	l := len(s)
	if l < n {
		return s
	}

	result := make([]byte, l)
	result[0] = s[0]

	maxOffset := (n - 1) * 2
	countList := make(map[int]int32)

	// 计算第一排的个数
	countList[0] = int32(math.Ceil(float64(l) / float64(maxOffset)))

	// 计算中间每行的个数
	for i := 1; i < n-1; i++ {
		countList[i] = 1
		tl, offset := l-i-1, (n-1-i)*2
		for tl > 0 {
			if tl >= offset {
				countList[i]++
			}
			tl -= offset
			offset = maxOffset - offset
		}
	}

	// 计算最后一排的个数
	//countList[n-1] = int32(math.Ceil(float64(l-n+1) / float64(maxOffset)))
	fmt.Println(countList)

	index := int32(0)
	addOffset, subOffset := int32(0), int32(0)
	j, ifAdd := 0, true
	for total := 1; total < l; total++ {
		if ifAdd {
			offset := countList[j]
			if j == 0 {
				offset = countList[j] + addOffset
			} else if j == n-2 {
				offset = countList[j] - addOffset
			}

			index += offset
			j++
			if j == n-1 {
				j--
				if n != 2 {
					addOffset++
				}
				ifAdd = false
			}
		} else {
			offset := countList[j]
			if j == n-2 {
				offset = countList[j] - 1 - subOffset
			} else if j == 0 {
				offset = countList[j] + subOffset
			}

			index -= offset
			j--
			if j < 0 {
				j++
				if n != 2 {
					subOffset++
				}
				ifAdd = true
			}
		}

		result[total] = s[index]
	}

	return string(result)
}
