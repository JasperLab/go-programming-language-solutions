package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(uint64(i))
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(uint64(i))
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(uint64(i))
	}
}

func TestPopCountShift(t *testing.T) {
	zero := uint64(0b0000000000000000000000000000000000000000000000000000000000000000)
	one := uint64(0b0000000000000000000000000000000000000000000000000000000000010000)
	two := uint64(0b0010000000000000000000000000000000000000000000000000000000000010)
	three := uint64(0b0000000000000000000000000000000000000000000000000010000000100010)
	all := uint64(0b1111111111111111111111111111111111111111111111111111111111111111)
	if PopCountShift(zero) != 0 {
		t.Errorf("PopCountShift(zero) = %d", PopCountShift(zero))
	}
	if PopCountShift(one) != 1 {
		t.Errorf("PopCountShift(one) = %d", PopCountShift(one))
	}
	if PopCountShift(two) != 2 {
		t.Errorf("PopCountShift(two) = %d", PopCountShift(two))
	}
	if PopCountShift(three) != 3 {
		t.Errorf("PopCountShift(three) = %d", PopCountShift(three))
	}
	if PopCountShift(all) != 64 {
		t.Errorf("PopCountShift(all) = %d", PopCountShift(all))
	}
}

func TestPopCountClear(t *testing.T) {
	zero := uint64(0b0000000000000000000000000000000000000000000000000000000000000000)
	one := uint64(0b0000000000000000000000000000000000000000000000000000000000010000)
	two := uint64(0b0010000000000000000000000000000000000000000000000000000000000010)
	three := uint64(0b0000000000000000000000000000000000000000000000000010000000100010)
	all := uint64(0b1111111111111111111111111111111111111111111111111111111111111111)
	if PopCountClear(zero) != 0 {
		t.Errorf("PopCountClear(zero) = %d", PopCountClear(zero))
	}
	if PopCountClear(one) != 1 {
		t.Errorf("PopCountClear(one) = %d", PopCountClear(one))
	}
	if PopCountClear(two) != 2 {
		t.Errorf("PopCountClear(two) = %d", PopCountClear(two))
	}
	if PopCountClear(three) != 3 {
		t.Errorf("PopCountClear(three) = %d", PopCountClear(three))
	}
	if PopCountClear(all) != 64 {
		t.Errorf("PopCountClear(all) = %d", PopCountClear(all))
	}
}
