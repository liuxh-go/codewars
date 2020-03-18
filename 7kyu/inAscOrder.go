package _kyu

import "sort"

/*
	判断输入的数组是否是升序排序的
*/

func InAscOrder(numbers []int) bool {
	return sort.IntsAreSorted(numbers)
}
