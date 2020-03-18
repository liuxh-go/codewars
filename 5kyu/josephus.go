package _kyu

/*
	给定一个数组和k,每隔k个位置选出一个数,直到数组被选完

	ex:
	[1,2,3,4,5,6,7] - initial sequence
	[1,2,4,5,6,7] => 3 is counted out and goes into the result [3]
	[1,2,4,5,7] => 6 is counted out and goes into the result [3,6]
	[1,4,5,7] => 2 is counted out and goes into the result [3,6,2]
	[1,4,5] => 7 is counted out and goes into the result [3,6,2,7]
	[1,4] => 5 is counted out and goes into the result [3,6,2,7,5]
	[4] => 1 is counted out and goes into the result [3,6,2,7,5,1]
	[] => 4 is counted out and goes into the result [3,6,2,7,5,1,4]
*/

func Josephus(items []interface{}, k int) []interface{} {
	l := len(items)
	result := make([]interface{}, 0, l)

	index := -1
	countedIndexMap := make(map[int]interface{})
	for n := 0; n < l; n++ {
		i, count := index, 0
		for {
			i = (i + 1) % l
			if _, exists := countedIndexMap[i]; exists {
				continue
			}

			count++
			if count == k {
				break
			}
		}

		result = append(result, items[i])
		countedIndexMap[i] = struct{}{}
		index = i
	}

	return result
}
