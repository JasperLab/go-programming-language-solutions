package main

import "testing"

func TestEmptyRotate(t *testing.T) {
	var a []int
	if a = rotate(a, 0); len(a) > 0 {
		t.Errorf("'[]' != %q", a)
	}
}

func TestSingleRotate(t *testing.T) {
	a := []int{1}
	if a = rotate(a, 0); len(a) != 1 || a[0] != 1 {
		t.Errorf("'[1]' != %q", a)
	}
}

func TestTwoRotate(t *testing.T) {
	a := []int{1,2}
	if a = rotate(a, 1); len(a) != 2 || a[0] != 2 || a[1] != 1 {
		t.Errorf("[2,1] != %v", a)
	}
}

func TestFiveRotate(t *testing.T) {
	a := []int{0,1,2,3,4}
	if a = rotate(a, 2); len (a) != 5 || a[0] != 2 || a[1] != 3 || a[2] != 4 || a[3] != 0 || a[4] != 1 {
		t.Errorf("[2,3,4,0,1] != %v", a)
	}
}

func TestOutOfRangeRotate(t *testing.T) {
	a := []int{1,2,3}
	if a = rotate(a, 5); len(a) != 3 || a[0] != 1 || a[1] != 2 {
		t.Errorf("[1,2,3] != %v", a)
	}
}

func TestNegativeRotate(t *testing.T) {
	a := []int{1,2,3}
	if a = rotate(a, -2); len(a) != 3 || a[0] != 1 || a[1] != 2 {
		t.Errorf("[1,2,3] != %v", a)
	}
}
