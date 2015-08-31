package solution

import "testing"

func TestBasic(t *testing.T) {
	exString := "()"
	if Solution(exString) != 1 {
		t.Errorf("%#v should be classified as 0", exString)
	}
}

func TestBasicWrong(t *testing.T) {
	exString := "([)[]]"
	if Solution(exString) != 0 {
		t.Errorf("%#v should be classified as 1", exString)
	}
}

func TestSolution(t *testing.T) {
	exString := "({[]}{[]})"
	if Solution(exString) != 1 {
		t.Errorf("%#v should be classified as 0", exString)
	}
}
