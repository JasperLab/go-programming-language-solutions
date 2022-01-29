package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("vova")
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	var v bytes.Buffer

	h := len(s) % 3
	v.WriteString(s[:h])
	v.WriteByte(',')

	for i := h; i < len(s); i += 3 {
		if i+3 >= len(s) {
			v.WriteString(s[i:])
		} else {
			v.WriteString(s[i : i+3])
			v.WriteByte(',')
		}
	}

	return v.String()
}
