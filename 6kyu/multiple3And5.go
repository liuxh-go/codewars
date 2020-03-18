package _kyu

/*
找到小于number的所有3和5的倍数的和
*/

func Multiple3And5(number int) int {
	result := 0

	for i := 3; i < number; i += 3 {
		result += i
	}

	for j, k := 5, 0; j < number; j += 5 {
		if k == 2 {
			k = 0
			continue
		}

		result += j
		k++
	}

	return result
}
