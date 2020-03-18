package _kyu

import (
	"math"
)

/*
	拆分n的平方和序列,必须要递增,尽可能大
*/

func Decompose(n int64) []int64 {
	result := make([]int64, 0)

	if dfs2(n*n, n, &result) {
		return result
	} else {
		return nil
	}
}

func dfs2(area, i int64, result *[]int64) bool {
	if area > 0 {
		for num := calc(area); num > 0 && num <= i; num-- {
			*result = append([]int64{num}, (*result)...)
			if num < i && dfs2(area-num*num, num, result) {
				return true
			}

			if len(*result) > 0 {
				*result = (*result)[1:]
			}
		}
	}

	return area == 0
}

func calc(area int64) int64 {
	return int64(math.Sqrt(float64(area)))
}
