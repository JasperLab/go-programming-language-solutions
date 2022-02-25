package main

import "testing"

var files [3]string = [3]string{"./empty.txt", "./single_word.txt", "./multiple_words.txt"}

func TestEmptyFile(t *testing.T) {
	e := sorted_counts_str(make(map[string]uint))
	a := sorted_counts_str(count_words(files[0:1]))
	if a != e {
		t.Errorf("%v != %v", a, e)
	}
}

func TestSingleWordFile(t *testing.T) {
	e := map[string]uint{"hello": 3}
	a := count_words(files[1:2])
	if sorted_counts_str(e) != sorted_counts_str(a) {
		t.Errorf("%v != %v", a, e)
	}
}

func TestMutlipleWordFile(t *testing.T) {
	e := map[string]uint{
		"Hello": 1,
		"мир":   1,
		"งเ":    1,
	}
	a := count_words(files[2:])
	if sorted_counts_str(e) != sorted_counts_str(a) {
		t.Errorf("%v != %v", a, e)
	}
}

func TestMultipleFiles(t *testing.T) {
	e := map[string]uint{
		"Hello": 1,
		"hello": 3,
		"мир":   1,
		"งเ":    1,
	}
	a := count_words(files[0:])
	if sorted_counts_str(e) != sorted_counts_str(a) {
		t.Errorf("%v != %v", a, e)
	}
}
