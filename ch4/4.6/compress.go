package main

import (
	"unicode"
	"unicode/utf8"
)

func compress(s []byte) []byte {
	if len(s) == 0 {
		return s
	}
	result := s[:0]
	for i, j := 0, 0; i < len(s); j++ {
		r, l := utf8.DecodeRune(s[i:]) 
		if unicode.IsSpace(r) {
			result = append(result, ' ')
			i += l
			continue
		}
		result = append(result, s[i])
		i++
	}
	return result 
}
