package _kyu

import "testing"

func TestMultiple3And5(t *testing.T) {
	t.Log(Multiple3And5(10))
}

func TestComp(t *testing.T) {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
	t.Log(Comp(a1, a2))
}

func TestCamelCase(t *testing.T) {
	t.Log(CamelCase("asfdfd adaf ffa s"))
}

func TestSquaresInRect(t *testing.T) {
	t.Log(SquaresInRect(20, 14))
}

func TestInterpreter(t *testing.T) {
	t.Log(Interpreter("++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++.+++++++.++++++++++.+++++++++.++++.++++++++.++++++++++++++++++++++++++++++++++++++.++++++.++++++..+++++.++++++++++++++++++++++++++.+++++++++++++++.+++++++++.+++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++.++++++++++++.+++++++++++++++++++.++++++.++.+++++++.+++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++.++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++.+++++++++++++++++.++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++..+++++++++++++.+++++.+++++++++++++++++++...++++++++++++++++++++++++++.+++.++++++++++++++++++++++++++++++++++++++++++.+++.++++++++++++++++.+++++++++++++++.++++++++++++++++++++.++++++++.++++++++++.+++++++++++++++++++++++++++++++.+++++++++.+++++++++++++++++++.+++++++++.+++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++.+++++++++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++.+++++++++++++.+++++++++++++++++++++++++++++++++++++++.++.++++++++++++++++++++.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++.+++++++++++++++++++++++.+++++++++.++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++"))
}

func TestWave(t *testing.T) {
	t.Log(Wave(""))
}
