package main

import (
	"bytes"
	"testing"
)

func TestEmptyReverse(t *testing.T) {
	s := []byte("")
	e := []byte("")
	if s = reverse(s); len(s) != 0 {
		t.Errorf("%v != %v", e, s)
	}
}

func TestSingleReverse(t *testing.T) {
	s := []byte("H")
	e := []byte("H")
	if s = reverse(s); !bytes.Equal(e, s) {
		t.Errorf("%v != %v", e, s)
	}
}

func TestSymmetricalReverse(t *testing.T) {
	s := []byte("\xe4\xb8\x96H\xe7\x95\x8c")
	e := []byte("\xe7\x95\x8cH\xe4\xb8\x96")
	if s = reverse(s); !bytes.Equal(s, e) {
		t.Errorf("%v != %v", e, s)
	}
}

func TestAssymetricalReverse(t *testing.T) {
	s := []byte("H\xe7\x95\x8c")
	e := []byte("\xe7\x95\x8cH")
	if s = reverse(s); !bytes.Equal(s, e) {
		t.Errorf("%v != %v", e, s)
	}
}

func TestComplexReverse(t *testing.T) {
	s := []byte("Hello, \xe4\xb8\x96\xe7\x95\x8c")
	e := []byte("\xe7\x95\x8c\xe4\xb8\x96 ,olleH")
	if s = reverse(s); !bytes.Equal(s, e) {
		t.Errorf("%v != %v", e, s)
	}
}
