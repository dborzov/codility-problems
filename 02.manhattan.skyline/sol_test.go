package solution

import (
	"fmt"
	"testing"
)

func TestOneLevel(t *testing.T) {
	exOneLev := []int{10}
	if Solution(exOneLev) != 1 {
		t.Errorf("one-level %v should have one level only", exOneLev)
	}

}

func Test3Levels(t *testing.T) {
	exH := []int{10, 20, 30, 40, 50, 60}
	if Solution(exH) != 6 {
		t.Errorf("ladder-shaped %v fails", exH)
	}

}

func TestTrimmer(t *testing.T) {
	tSlice := []int{2, 3, 4, 5}
	tLev := 4
	tSlice, res := isIn(tSlice, tLev)
	fmt.Printf("We get %#v, %#v \n", res, tSlice)

}
