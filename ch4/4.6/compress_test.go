package main

import (
	"bytes"
	"testing"
)

func TestEmpty(t *testing.T) {
	var s []byte
	if s = compress(s); len(s) != 0 {
		t.Errorf("[] != %v", s)
	}
}

func TestSingleNonSpace(t *testing.T) {
	s := []byte("V")
	if s = compress(s); len(s) != 1 || s[0] != 'V' {
		t.Errorf("[V] != %v", s)
	}
}

func TestSingleSpace(t *testing.T) {
	s := []byte("V V")
	if s = compress(s); len(s) != 3 || s[0] != 'V' || s[1] != ' ' || s[2] != 'V' {
		t.Errorf("[V V] != %v", s)
	}
}

func TestSingeUtf(t *testing.T) {
	s := []byte("\u00a0")
	if s = compress(s); len(s) != 1 || s[0] != ' ' {
		t.Errorf("[ ] != %v", s)
	}
}

func TestMultiples(t *testing.T) {
	s := []byte("1 2\u00a03\u00854")
	e := []byte("1 2 3 4")
	if s = compress(s); !bytes.Equal(s, e) {
		t.Errorf("%v != %v", e, s)
	}
}


