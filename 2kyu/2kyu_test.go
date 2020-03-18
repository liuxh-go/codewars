package _kyu

import (
	"fmt"
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	clues := []int{
		0, 3, 0, 5, 3, 4,
		0, 0, 0, 0, 0, 1,
		0, 3, 0, 3, 2, 3,
		3, 2, 0, 3, 1, 0,
	}
	result := SolvePuzzle(clues)

	for _, items := range result {
		fmt.Println(items)
	}
}
