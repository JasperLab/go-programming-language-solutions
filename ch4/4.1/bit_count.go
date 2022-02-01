package main


func numOfBits(l, r [32]byte) uint8 {
	var sum uint8
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			lv := l[i] >> j 
			rv := r[i] >> j
			if lv & 1 != rv & 1 {
				sum++
			}
		}
	}
	return sum
}
