package solution

import "testing"

func TestEmpty(t *testing.T) {
	if Solution("") != 1 {
		t.Errorf("empty string should return OK")
	}
}

func TestOneBracket(t *testing.T) {
	if Solution("(()())(") != 0 {
		t.Errorf("broken string is not reported as broken")
	}
}
