package main

import(
	"testing"
)

func testZeroReverse(t *testing.T) {
	a := [...]int{0, 0, 0, 0, 0, 0}
	e := [...]int{0, 0, 0, 0, 0, 0}
	if reverse(&a); a != e {
		t.Errorf("%q != %q", e, a)
	}
}

func testOrderReverse(t *testing.T) {
	a := [...]int{0,1,2,3,4,5}
	e := [...]int{5,4,3,2,1,0}
	if reverse(&a); a != e {
		t.Errorf("[5,4,3,2,1,0] != %q", a)
	}
}

