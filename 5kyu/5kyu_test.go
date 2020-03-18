package _kyu

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	t.Log(FirstNonRepeating("sTreSS"))
}

func TestPerimeter(t *testing.T) {
	t.Log(Perimeter(31))
}

func TestAlphanumeric(t *testing.T) {
	t.Log(Alphanumeric("hello world_"))
}

func TestProductFib(t *testing.T) {
	t.Log(ProductFib(8))
}

func TestCreateSpiral(t *testing.T) {
	result := CreateSpiral(0)

	for _, items := range result {
		t.Log(items)
	}
}

func TestJosephus(t *testing.T) {
	t.Log(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3))
}
