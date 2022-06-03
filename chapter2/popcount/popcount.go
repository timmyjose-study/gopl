package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)

	}
}

func PopCount1(n uint64) int {
	return int(pc[byte(n>>0)] +
		pc[byte(n>>8)] +
		pc[byte(n>>16)] +
		pc[byte(n>>24)] +
		pc[byte(n>>32)] +
		pc[byte(n>>40)] +
		pc[byte(n>>48)] +
		pc[byte(n>>64)])
}

func PopCount2(n uint64) int {
	count := 0

	for i := 0; i < 8; i++ {
		count += int(pc[byte(n>>(i*8))])
	}

	return count
}

func PopCount3(n uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((n >> i) & 0x1)
	}

	return count
}

func PopCount4(n uint64) int {
	count := 0
	for n != 0 {
		count++
		n &= n - 1
	}

	return count
}