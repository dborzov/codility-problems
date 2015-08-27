package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	// example := []int{10}

}

func TestTrimmer(t *testing.T) {
	tSlice := []int{2, 3, 4, 5}
	tLev := 4
	tSlice, res := isIn(tSlice, tLev)
	fmt.Printf("We get %#v, %#v \n", res, tSlice)

}
