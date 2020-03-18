package _kyu

import (
	"fmt"
	"sort"
)

/*
	计算输入数组中的质数

	输出单个格式(质数 该质数在数组中能被整除的数的和)

	格式:
	(2 12)(3 27)(5 15)
*/

func SumOfDivided(lst []int) string {
	primeMap := make(map[int]int)
	primeSlice := make([]int, 0)
	for _, num := range lst {
		// 计算该数字的所有质数因子
		primeList := primeList(num)
		for _, prime := range primeList {
			// 记录该质数因子关联的所有数字的和
			if _, exists := primeMap[prime]; exists == false {
				primeSlice = append(primeSlice, prime)
			}
			primeMap[prime] += num
		}
	}

	// 组装输出
	sort.Ints(primeSlice)
	result := ""
	for _, prime := range primeSlice {
		result += fmt.Sprintf("(%d %d)", prime, primeMap[prime])
	}

	return result
}

// 求一个数的质数因子列表
func primeList(n int) []int {
	result := make([]int, 0)

	if n < 0 {
		n = -n
	}

	prime := 0
	for i := 2; i <= n; {
		if n%i == 0 {
			if prime != i {
				prime = i
				result = append(result, i)
			}

			n /= i
		} else {
			i++
		}
	}

	return result
}
