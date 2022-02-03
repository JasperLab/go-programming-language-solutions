package main

func reverse(p *[6]int) {
	for i, j := 0, 5; i < j; i, j = i + 1, j - 1 {
		p[i], p[j] = p[j], p[i]
	}
}
