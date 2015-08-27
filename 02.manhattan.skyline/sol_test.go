package solution

import "testing"

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

func TestBump(t *testing.T) {
	exH := []int{20, 20, 30, 30, 10, 10, 10}
	if Solution(exH) != 3 {
		t.Errorf("bump-shaped %v fails", exH)
	}

}

func TestPyramid(t *testing.T) {
	exH := []int{20, 20, 30, 30, 20, 20, 20}
	if Solution(exH) != 2 {
		t.Errorf("pyramid-shaped %v fails", exH)
	}

}

func TestPit(t *testing.T) {
	exH := []int{20, 20, 10, 10, 20, 20, 20}
	if Solution(exH) != 3 {
		t.Errorf("pit-shaped %v fails", exH)
	}

}
