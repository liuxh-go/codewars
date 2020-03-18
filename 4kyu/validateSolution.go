package _kyu

/*
	校验数独
*/

func ValidateSolution(m [][]int) bool {
	// 校验行
	for i := 0; i < 9; i++ {
		checkList := make([]int, 9)
		for j := 0; j < 9; j++ {
			if m[i][j] == 0 {
				return false
			}

			if checkList[m[i][j]-1] == 1 {
				return false
			}

			checkList[m[i][j]-1] = 1
		}
	}

	// 校验列
	for j := 0; j < 9; j++ {
		checkList := make([]int, 9)
		for i := 0; i < 9; i++ {
			if m[i][j] == 0 {
				return false
			}

			if checkList[m[i][j]-1] == 1 {
				return false
			}

			checkList[m[i][j]-1] = 1
		}
	}

	// 校验区
	// 转换区 0 - 9区
	// i/3 + (j/3) * 3
	checkMap := make(map[int][]int)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if m[i][j] == 0 {
				return false
			}

			area := j/3 + (i/3)*3

			if checkList, exists := checkMap[area]; exists && checkList[m[i][j]-1] == 1 {
				return false
			}

			if _, exists := checkMap[area]; exists == false {
				checkMap[area] = make([]int, 9)
			}

			checkMap[area][m[i][j]-1] = 1
		}
	}

	return true
}
