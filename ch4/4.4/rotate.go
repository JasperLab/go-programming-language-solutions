package main

func rotate(a []int, pos int) []int {
	if pos < 0 {
		return a
	}
	if pos >= len(a) {
		return a
	}
	b := make([]int, len(a))
	j := pos
	for i := 0; i < len(a); i++ {
		b[i] = a[j]
		j++
		if j == len(a) {
			j = 0
		}
	}
	return b
}
