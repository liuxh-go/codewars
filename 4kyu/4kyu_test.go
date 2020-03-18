package _kyu

import (
	"testing"
)

func TestMix(t *testing.T) {
	t.Log(Mix("looping is fun but dangerous", "less dangerous than coding"))
}

func TestSolvePuzzle(t *testing.T) {
	in := []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 0,
	}

	out := SolvePuzzle(in)
	for _, o := range out {
		t.Log(o)
	}
}

func TestDblLinear(t *testing.T) {
	for i := 500; i < 501; i++ {
		t.Log(DblLinear(i))
	}
}

func TestSolution(t *testing.T) {
	t.Log(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}

func TestBase91(t *testing.T) {
	Encode([]byte("t"))
}

func TestSumOfDivided(t *testing.T) {
	t.Log(SumOfDivided([]int{15, 21, 24, 30, 45}))
}

func TestDecompose(t *testing.T) {
	t.Log(Decompose(50))
}

func TestValidateSolution(t *testing.T) {
	sudoku := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	t.Log(ValidateSolution(sudoku))
}
