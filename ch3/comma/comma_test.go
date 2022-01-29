package main

import "testing"

func TestEmptyString(t *testing.T) {
	r := comma("")
	if r != "" {
		t.Errorf("%q != %q", r, "")
	}
}

func TestSingleDigit(t *testing.T) {
	if r := comma("1"); r != "1" {
		t.Errorf("%q != %q", r, "1")
	}
}

func TestThreeDigits(t *testing.T) {
	if r := comma("123"); r != "123" {
		t.Errorf("%q != %q", r, "123")
	}
}

func TestTenDigits(t *testing.T) {
	if r := comma("1234567890"); r != "1,234,567,890" {
		t.Errorf("%q != %q", r, "1,234,567,890")
	}
}

func TestWithLetters(t *testing.T) {
	if r := comma("abcd"); r != "a,bcd" {
		t.Errorf("%q != %q", r, "a,bcd")
	}
}
