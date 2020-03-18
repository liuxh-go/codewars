package _kyu

import (
	"fmt"
	"sort"
)

var (
	clueNodeMap = map[int]map[int][][]int{}
)

const (
	N = 4
)

type selfClue struct {
	index        int
	forward, rev int
}

/*
	更换思路:

	递归规则,直接填一列或一行,若没法填下一个规则,则返回上一个规则,当规则填完时,就返回
*/

func init() {
	// 生成全排列列表
	data := make([]int, N)
	for i := 0; i < N; i++ {
		clueNodeMap[i+1] = make(map[int][][]int)
		data[i] = i + 1
	}
	result := PermutationConcurrency(data)

	for _, items := range result {
		// 统计正向反向数出的层数
		count := calcCount(items)
		reCount := calcCount(reverse(items))

		// 组装规则字符串字典
		clueNodeMap[count][reCount] = append(clueNodeMap[count][reCount], items)
		clueNodeMap[count][0] = append(clueNodeMap[count][0], items)
	}
}

func SolvePuzzle(clues []int) [][]int {
	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, N)
	}

	selfClueList := make([]*selfClue, 0)
	existMap := make(map[int]interface{})
	for i := 0; i < len(clues); i++ {
		if clues[i] == 0 {
			continue
		}

		// 组装有限制的规则列表
		if i < N {
			selfClueList = append(selfClueList, &selfClue{i, clues[i], clues[3*N-1-i]})

			if clues[3*N-1-i] != 0 {
				existMap[3*N-1-i] = struct{}{}
			}
		} else if i < 2*N {
			selfClueList = append(selfClueList, &selfClue{i, clues[i], clues[5*N-1-i]})

			if clues[5*N-1-i] != 0 {
				existMap[5*N-1-i] = struct{}{}
			}
		} else {
			// 需要去掉已经在双向列表中的规则
			if _, exists := existMap[i]; exists {
				continue
			}

			selfClueList = append(selfClueList, &selfClue{i, clues[i], 0})
		}
	}

	sort.Slice(selfClueList, func(i, j int) bool {
		// 排序规则
		// 1、forward和rev都不为0的排前面
		// 2、两个clue中forward, rev最大的排前面
		maxNum := 0
		if selfClueList[j].forward > maxNum {
			maxNum = selfClueList[j].forward
		}

		if selfClueList[j].rev > maxNum {
			maxNum = selfClueList[j].rev
		}

		if selfClueList[i].rev == 0 && selfClueList[j].rev != 0 {
			return false
		}

		if selfClueList[j].rev == 0 && selfClueList[i].rev != 0 {
			return true
		}

		return selfClueList[i].forward > maxNum || selfClueList[i].rev > maxNum
	})

	clueIndexList := make([]int, 4*N)
	if dfs(0, len(selfClueList), result, selfClueList, clueIndexList) == false {
		fmt.Println("failed!")
	}

	if dfs1(0, result) == false {
		fmt.Println("failed1!")
	}

	return result
}

func dfs(level, clueLevel int, result [][]int, cluesList []*selfClue, clueIndexList []int) bool {
	if level > clueLevel-1 {
		return true
	}

	clues := cluesList[level]
	if clues.index < N {
		// 从上到下,列
		cluesNodeList := clueNodeMap[clues.forward][clues.rev]
		for cluesListIndex := clueIndexList[level]; cluesListIndex < len(cluesNodeList); cluesListIndex++ {
			fillData := cluesNodeList[cluesListIndex]

			nowData := make([]int, N)
			for i := 0; i < N; i++ {
				nowData[i] = result[i][clues.index]
			}

			if isMatch(fillData, nowData, result, false, clues.index) {
				clueIndexList[level] = cluesListIndex + 1
				fillColumn(result, fillData, clues.index)
				if dfs(level+1, clueLevel, result, cluesList, clueIndexList) {
					return true
				}
				clueIndexList[level] = 0
				fillColumn(result, nowData, clues.index)
			}
		}
	} else if clues.index < 2*N {
		// 反向,行,需要将数据倒转然后从左到右
		cluesNodeList := clueNodeMap[clues.forward][clues.rev]
		for cluesListIndex := clueIndexList[level]; cluesListIndex < len(cluesNodeList); cluesListIndex++ {
			fillData := reverse(cluesNodeList[cluesListIndex])

			nowData := make([]int, N)
			for j := 0; j < N; j++ {
				nowData[j] = result[clues.index-N][j]
			}

			if isMatch(fillData, nowData, result, true, clues.index-N) {
				clueIndexList[level] = cluesListIndex + 1
				fillRow(result, fillData, clues.index-N)
				if dfs(level+1, clueLevel, result, cluesList, clueIndexList) {
					return true
				}
				clueIndexList[level] = 0
				fillRow(result, nowData, clues.index-N)
			}
		}
	} else if clues.index < 3*N {
		// 反向,列,需要将数据倒转然后从上到下
		cluesNodeList := clueNodeMap[clues.forward][clues.rev]
		for cluesListIndex := clueIndexList[level]; cluesListIndex < len(cluesNodeList); cluesListIndex++ {
			fillData := reverse(cluesNodeList[cluesListIndex])

			nowData := make([]int, N)
			for i := 0; i < N; i++ {
				nowData[i] = result[i][3*N-1-clues.index]
			}

			if isMatch(fillData, nowData, result, false, 3*N-1-clues.index) {
				clueIndexList[level] = cluesListIndex + 1
				fillColumn(result, fillData, 3*N-1-clues.index)
				if dfs(level+1, clueLevel, result, cluesList, clueIndexList) {
					return true
				}
				clueIndexList[level] = 0
				fillColumn(result, nowData, 3*N-1-clues.index)
			}
		}
	} else {
		// 从左到右,行
		cluesNodeList := clueNodeMap[clues.forward][clues.rev]
		for cluesListIndex := clueIndexList[level]; cluesListIndex < len(cluesNodeList); cluesListIndex++ {
			fillData := cluesNodeList[cluesListIndex]

			nowData := make([]int, N)
			for j := 0; j < N; j++ {
				nowData[j] = result[4*N-1-clues.index][j]
			}

			if isMatch(fillData, nowData, result, true, 4*N-1-clues.index) {
				clueIndexList[level] = cluesListIndex + 1
				fillRow(result, fillData, 4*N-1-clues.index)
				if dfs(level+1, clueLevel, result, cluesList, clueIndexList) {
					return true
				}
				clueIndexList[level] = 0
				fillRow(result, nowData, 4*N-1-clues.index)
			}
		}
	}

	return false
}

func dfs1(n int, result [][]int) bool {
	if n > N*N-1 {
		return true
	}

	i, j := n/N, n%N
	if result[i][j] != 0 {
		return dfs1(n+1, result)
	}

	for num := 1; num <= N; num++ {
		if isValid(result, i, j, num) {
			result[i][j] = num
			if dfs1(n+1, result) {
				return true
			}
			result[i][j] = 0
		}
	}

	return false
}

// 校验行列是否重复
func isValid(result [][]int, i, j, num int) bool {
	for k := 0; k < N; k++ {
		if k != j && result[i][k] != 0 && result[i][k] == num {
			return false
		}
	}
	for k := 0; k < N; k++ {
		if k != i && result[k][j] != 0 && result[k][j] == num {
			return false
		}
	}

	return true
}

// 填充行
func fillRow(result [][]int, data []int, row int) {
	for j := 0; j < N; j++ {
		result[row][j] = data[j]
	}
}

// 填充列
func fillColumn(result [][]int, data []int, column int) {
	for i := 0; i < N; i++ {
		result[i][column] = data[i]
	}
}

func reverse(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func calcCount(data []int) int {
	start, count := data[0], 1
	for _, n := range data[1:] {
		if n > start {
			count++
			start = n
		}
	}

	return count
}

func isMatch(data []int, existsNum []int, result [][]int, isRow bool, index int) bool {
	for i := 0; i < N; i++ {
		if existsNum[i] == 0 {
			continue
		}

		if existsNum[i] != data[i] {
			return false
		}
	}

	if isRow {
		// 行需要判断每一列是否有重复
		for j := 0; j < N; j++ {
			if existsNum[j] != 0 {
				continue
			}

			for i := 0; i < N; i++ {
				if i == index {
					continue
				}

				if result[i][j] != 0 && result[i][j] == data[j] {
					return false
				}
			}
		}
	} else {
		// 列需要判断每一行是否有重复
		for i := 0; i < N; i++ {
			if existsNum[i] != 0 {
				continue
			}

			for j := 0; j < N; j++ {
				if j == index {
					continue
				}

				if result[i][j] != 0 && result[i][j] == data[i] {
					return false
				}
			}
		}
	}

	return true
}

func prefixIncrement(in []int, s []int, next chan []int) {
	for _, c := range s {
		exist := false
		for _, e := range in {
			if e == c {
				exist = true
				break
			}
		}
		if exist {
			continue
		}

		temp := make([]int, 0)
		temp = append(temp, in...)
		temp = append(temp, c)
		next <- temp
	}
}

func permutaionConImpl(req chan []int, out chan []int, s []int) {
	go func() {
		//递归退出条件: len(v) == len(s)-1
		v, ok := <-req
		if !ok {
			return
		}

		next := out
		if len(v) != len(s)-1 {
			next = make(chan []int)
			permutaionConImpl(next, out, s)
		}

		prefixIncrement(v, s, next)
		for in := range req {
			prefixIncrement(in, s, next)
		}
		close(next)
	}()
}

// PermutationConcurrency  并发计算全排列
func PermutationConcurrency(s []int) [][]int {
	req, out := make(chan []int), make(chan []int)

	//开启goroutine计算
	permutaionConImpl(req, out, s)

	over := make(chan [][]int)

	//要开goroutine读取out，如果放在主函数中，会导致死锁。
	go func() {
		result := make([][]int, 0)

		for res := range out {
			result = append(result, res)
		}

		over <- result
	}()

	for _, c := range s {
		sl := []int{c}
		req <- sl
	}
	close(req)

	return <-over
}
