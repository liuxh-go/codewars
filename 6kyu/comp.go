package _kyu

import "math"

/*
	array2中所有的数的平方根是否都能在array1中找到,若都能找到则返回true;否则返回false

	array1或array2为空返回false
*/

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || len(array1) == 0 || len(array2) == 0 {
		return false
	}

	existsMap := make(map[int]interface{})

	for _, num2 := range array2 {
		if _, ok := existsMap[num2]; ok {
			continue
		}

		sq := math.Sqrt(float64(num2))
		if float64(int(sq)) != sq {
			return false
		}

		exists := false
		for _, num1 := range array1 {
			if num1 == int(sq) {
				exists = true
				break
			}
		}

		if exists == false {
			return false
		}

		existsMap[num2] = struct{}{}
	}

	return true
}
