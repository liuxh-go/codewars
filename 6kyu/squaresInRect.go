package _kyu

/*
	给定一个矩形的长和宽,拆分矩形为多个正方形
*/

func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}

	return squaresInRect(lng, wdth)
}

func squaresInRect(lng int, wdth int) []int {
	if lng < wdth {
		return append([]int{lng}, squaresInRect(lng, wdth-lng)...)
	} else if lng > wdth {
		return append([]int{wdth}, squaresInRect(wdth, lng-wdth)...)
	} else {
		return []int{wdth}
	}
}
