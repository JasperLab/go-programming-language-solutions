package main

import "testing"

func TestEmptyDedup(t *testing.T) {
	var a []string
	if a := dedup(a); len(a) != 0 {
		t.Errorf("[] != %v", a)
	}
}

func TestOneDedup(t *testing.T) {
	a := []string{"test"}
	if a := dedup(a); len(a) != 1 || a[0] != "test" {
		t.Errorf("[test] != %v", a)
	}
}

func TestSingleDedup(t *testing.T) {
	a := []string{"test", "test"}
	if a := dedup(a); len(a) != 1 || a[0] != "test" {
		t.Errorf("[test] != %v", a)
	}
}

func TestFewDedups(t *testing.T) {
	a := []string{"test1", "test1", "test2", "test3", "test3"}
	if a := dedup(a); len(a) != 3 || a[0] != "test1" || a[1] != "test2" || a[2] != "test3" {
		t.Errorf("[test1, test2, test3] != %v", a)
	}
}

