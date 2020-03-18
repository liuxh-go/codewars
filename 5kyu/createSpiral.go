package _kyu

/*
	给定一个数字n

	生成n*n的螺旋二维数组

	ex:
	n=3
	1	2	3
	8	9	4
	7	6	5

	n=4
	1   2   3   4
	12  13  14  5
	11  16  15  6
	10  9   8   7

	n=5
	1	2	3	4	5
	16	17	18	19	6
	15	24	25	20	7
	14	23	22	21	8
	13	12	11	10	9
*/

func CreateSpiral(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for i, j, start := 0, 0, 1; n > 0; i, j = i+1, j+1 {
		start = fillData(i, j, start, n, result)
		n -= 2
	}

	return result
}

func fillData(i, j, start, n int, result [][]int) int {
	if n == 1 {
		result[i][j] = start
		return start
	}

	row, column := i, j
	for ; column < j+n-1; column++ {
		result[i][column] = start
		start++
	}

	for ; row < n+i-1; row++ {
		result[row][column] = start
		start++
	}

	for ; column > j; column-- {
		result[row][column] = start
		start++
	}

	for ; row > i; row-- {
		result[row][column] = start
		start++
	}

	return start
}
