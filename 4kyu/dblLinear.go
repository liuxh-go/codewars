package _kyu

import (
	"math"
	"sort"
)

/*
	返回一个序列f(0)=1
	x => 2*x+1 和 3*x+1 也必须在序列中
	Ex: u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]
	1 2x+1 3x+1 2(2x+1)+1 3(2x+1)+1 2(3x+1)+1 3(3x+1)+1...
*/

func DblLinear(n int) int {
	result := []int{1}

	num, exp, count := 0, 1, int(math.Log10(float64(n)))
	for ; num < n; count++ {
		num += 2 * exp
		exp *= 2
	}

	mul(&result, []int{1}, count)
	sort.Ints(result)

	index := 0
	for i, j := 0, 0; j < n; i++ {
		if result[i] != result[i+1] {
			j++
		}

		index++
	}

	return result[index]
}

func mul(result *[]int, numList []int, n int) []int {
	if n <= 0 {
		return []int{}
	}

	data := make([]int, 0, len(numList)*2)
	for _, num := range numList {
		data = append(data, 2*num+1, 3*num+1)
	}

	*result = append(*result, data...)

	return mul(result, data, n-1)
}
