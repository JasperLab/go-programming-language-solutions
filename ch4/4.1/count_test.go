package main

import(
	"crypto/sha256"
	"testing"
)

func TestNoDiff(t *testing.T) {
	l, r := sha256.Sum256([]byte("x")), sha256.Sum256([]byte("x")) 
	if d := numOfBits(l, r); d != 0 {
		t.Errorf("expected: %d\t received: %d", 0, d)
	}
}

func TestOneDiff(t *testing.T) {
	l := sha256.Sum256([]byte("x"))
	r := l 
	r[0] = r[0]&(r[0]-1)
	if d := numOfBits(l, r); d != 1 {
		t.Errorf("expected: %d\t received: %d", 1, d)
	}
}

func TestManyDiff(t *testing.T) {
	l := sha256.Sum256([]byte("x"))
	r := sha256.Sum256([]byte("X"))
	if d := numOfBits(l, r); d <= 1 {
		t.Errorf("expected: more than 1\treceived: %d", d)
	}
}
