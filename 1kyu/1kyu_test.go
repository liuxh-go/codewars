package _kyu

import (
	"fmt"
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	clues := []int{7, 0, 0, 0, 2, 2, 3, 0, 0, 3, 0, 0, 0, 0, 3, 0, 3, 0, 0, 5, 0, 0, 0, 0, 0, 5, 0, 4}

	result := SolvePuzzle(clues)

	for _, item := range result {
		for _, i := range item {
			fmt.Printf("%d, ", i)
		}

		fmt.Println()
	}
}

func TestG(t *testing.T) {
	clues := []int{
		0, 0, 5, 0, 0, 0, 6,
		4, 0, 0, 2, 0, 2, 0,
		0, 5, 2, 0, 0, 0, 5,
		0, 3, 0, 5, 0, 0, 3,
	}
	result := SolvePuzzle(clues)

	for _, items := range result {
		fmt.Println(items)
	}
}
