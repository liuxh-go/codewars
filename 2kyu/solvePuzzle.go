package _kyu

import (
	"fmt"
	"sort"
)

var (
	// key1:正向层数;key2:反向层数;value:数组
	clueListMap = map[int]map[int][][]int{}
)

const (
	N = 6
)

// 自定义规则节点(记录非0规则正向和反向层数以及索引)
type SelfClueNode struct {
	Index        int
	Forward, Rev int
}

/*
	更换思路:

	1、递归规则,直接填一列或一行,若没法填下一个规则,则返回上一个规则,当规则填完时,就返回
	2、规则填完后再填充为0的格子,若为0的格子不能正确填写,则继续返回1的递归中
*/

func init() {
	// 生成全排列列表
	data := make([]int, N)
	for i := 0; i < N; i++ {
		clueListMap[i+1] = make(map[int][][]int)
		data[i] = i + 1
	}
	result := PermutationConcurrency(data)

	for _, items := range result {
		// 统计正向反向数出的层数
		count := calcCount(items)
		reCount := calcCount(reverse(items))

		// 组装规则字符串字典
		clueListMap[count][reCount] = append(clueListMap[count][reCount], items)
		clueListMap[count][0] = append(clueListMap[count][0], items)
	}
}

func SolvePuzzle(clues []int) [][]int {
	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, N)
	}

	selfClueNodeList := make([]*SelfClueNode, 0)
	existMap := make(map[int]interface{})
	for i := 0; i < len(clues); i++ {
		if clues[i] == 0 {
			continue
		}

		// 组装有限制的规则列表
		if i < N {
			selfClueNodeList = append(selfClueNodeList, &SelfClueNode{i, clues[i], clues[3*N-1-i]})

			if clues[3*N-1-i] != 0 {
				existMap[3*N-1-i] = struct{}{}
			}
		} else if i < 2*N {
			selfClueNodeList = append(selfClueNodeList, &SelfClueNode{i, clues[i], clues[5*N-1-i]})

			if clues[5*N-1-i] != 0 {
				existMap[5*N-1-i] = struct{}{}
			}
		} else {
			// 需要去掉已经在双向列表中的规则
			if _, exists := existMap[i]; exists {
				continue
			}

			selfClueNodeList = append(selfClueNodeList, &SelfClueNode{i, clues[i], 0})
		}
	}

	sort.Slice(selfClueNodeList, func(i, j int) bool {
		// 排序规则
		// 1、forward和rev都不为0的排前面
		// 2、两个clue中forward, rev最大的排前面
		maxNum := 0
		if selfClueNodeList[j].Forward > maxNum {
			maxNum = selfClueNodeList[j].Forward
		}

		if selfClueNodeList[j].Rev > maxNum {
			maxNum = selfClueNodeList[j].Rev
		}

		if selfClueNodeList[i].Rev == 0 && selfClueNodeList[j].Rev != 0 {
			return false
		}

		if selfClueNodeList[j].Rev == 0 && selfClueNodeList[i].Rev != 0 {
			return true
		}

		return selfClueNodeList[i].Forward > maxNum || selfClueNodeList[i].Rev > maxNum
	})

	clueIndexList := make([]int, len(selfClueNodeList))
	if dfs(0, result, selfClueNodeList, clueIndexList) == false {
		fmt.Println("failed!")
	}

	return result
}

func dfs(level int, result [][]int, cluesList []*SelfClueNode, clueIndexList []int) bool {
	if level > len(cluesList)-1 {
		// 规则填完后,还需要校验为0的格子

		// 统计结果中为0的格子
		zeroNodeList := make([][2]int, 0)
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if result[i][j] == 0 {
					zeroNodeList = append(zeroNodeList, [2]int{i, j})
				}
			}
		}

		if len(zeroNodeList) == 0 {
			return true
		}

		return dfs1(0, result, zeroNodeList)
	}

	clues := cluesList[level]
	cluesNodeList := clueListMap[clues.Forward][clues.Rev]
	var isRow bool
	var curIndex int
	var ifReverse bool
	nowData := make([]int, N)

	if clues.Index < N {
		// 从上到下,列
		ifReverse = false
		isRow = false
		curIndex = clues.Index
		for i := 0; i < N; i++ {
			nowData[i] = result[i][curIndex]
		}
	} else if clues.Index < 2*N {
		// 反向,行,需要将数据倒转然后从左到右
		ifReverse = true
		isRow = true
		curIndex = clues.Index - N
		for j := 0; j < N; j++ {
			nowData[j] = result[curIndex][j]
		}
	} else if clues.Index < 3*N {
		// 反向,列,需要将数据倒转然后从上到下
		ifReverse = true
		isRow = false
		curIndex = 3*N - 1 - clues.Index
		for i := 0; i < N; i++ {
			nowData[i] = result[i][curIndex]
		}
	} else {
		// 从左到右,行
		ifReverse = false
		isRow = true
		curIndex = 4*N - 1 - clues.Index
		for j := 0; j < N; j++ {
			nowData[j] = result[curIndex][j]
		}
	}

	for cluesListIndex := clueIndexList[level]; cluesListIndex < len(cluesNodeList); cluesListIndex++ {
		data := cluesNodeList[cluesListIndex]
		if ifReverse {
			data = reverse(data)
		}
		if isMatch(data, nowData, result, isRow, curIndex) {
			clueIndexList[level] = cluesListIndex + 1
			fillData(result, data, curIndex, isRow)
			if dfs(level+1, result, cluesList, clueIndexList) {
				return true
			}
			clueIndexList[level] = 0
			fillData(result, nowData, curIndex, isRow)
		}
	}

	return false
}

func dfs1(level int, result [][]int, zeroNodeList [][2]int) bool {
	if level > len(zeroNodeList)-1 {
		return true
	}

	i, j := zeroNodeList[level][0], zeroNodeList[level][1]
	for num := 1; num <= N; num++ {
		if isValid(result, i, j, num) {
			result[i][j] = num
			if dfs1(level+1, result, zeroNodeList) {
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

// 填充数据
func fillData(result [][]int, data []int, index int, isRow bool) {
	if isRow {
		for j := 0; j < N; j++ {
			result[index][j] = data[j]
		}
	} else {
		for i := 0; i < N; i++ {
			result[i][index] = data[i]
		}
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
	max, count := data[0], 1
	for _, cur := range data[1:] {
		if cur > max {
			count++
			max = cur
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
