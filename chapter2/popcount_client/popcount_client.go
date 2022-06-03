package main

import (
	"fmt"
	"popcount"
)

func main() {
	var n uint64

	fmt.Scanf("%d", &n)

	fmt.Printf("Popcount1(%d) = %d\n", n, popcount.PopCount1(n))
	fmt.Printf("Popcount2(%d) = %d\n", n, popcount.PopCount2(n))
	fmt.Printf("Popcount3(%d) = %d\n", n, popcount.PopCount3(n))
	fmt.Printf("Popcount4(%d) = %d\n", n, popcount.PopCount4(n))
}