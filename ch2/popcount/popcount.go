package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var sum int
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func PopCountShift(x uint64) int {
	sum := int(x & 1)
	for i := 0; i < 63; i++ {
		x = x >> 1
		sum += int(x & 1)
	}
	return sum
}

func PopCountClear(x uint64) int {
	sum := 0

	for x != 0 {
		sum += 1
		x = x & (x - 1)
	}

	return sum
}
