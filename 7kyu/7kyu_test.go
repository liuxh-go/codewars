package _kyu

import "testing"

func TestInAscOrder(t *testing.T) {
	t.Log(InAscOrder([]int{1, 2, 3, 4, 5}))
	t.Log(InAscOrder([]int{5, 4, 3, 2, 1}))
}
