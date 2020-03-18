package _kyu

/*
	斐波那契数列
*/

func Perimeter(n int) int {
	a, b, result := 1, 1, 0
	for i := 0; i <= n; i++ {
		result += b
		a, b = a+b, a
	}

	return result * 4
}
