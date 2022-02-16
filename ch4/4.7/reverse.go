package main

import "unicode/utf8"

func reverse(s []byte) []byte {
	// first pass: reverse bytes
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	// second pass: fix multi-byte UTF-8 runes
	for i := 0; i < len(s); i++ {
		if utf8.Valid(s[i : i+1]) {
			continue
		}

		//rune not valid: rotate two bytes around
		s[i], s[i+1] = s[i+1], s[i]
		if utf8.Valid(s[i : i+2]) {
			i += 1
			continue
		}

		//still not valid: rotate three bytes
		s[i], s[i+1], s[i+2] = s[i+2], s[i], s[i+1]
		if utf8.Valid(s[i : i+3]) {
			i += 2
			continue
		}

		//still not valid: rotate four byte
		s[i], s[i+1], s[i+2], s[i+3] = s[i+3], s[i], s[i+1], s[i+2]
		i += 3
	}
	return s
}
