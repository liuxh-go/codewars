package _kyu

/*
	从斐波那契数列中找到两个相邻的数的乘积等于输入的数,若存在则返回s1,s2,1;不存在则返回s1,s2,0(s1*s2>prod)
*/

func ProductFib(prod uint64) [3]uint64 {
	var a, b uint64 = 2, 1
	for {
		a, b = a+b, a
		if a*b == prod {
			return [3]uint64{b, a, 1}
		}

		if a*b > prod {
			return [3]uint64{b, a, 0}
		}
	}
}
