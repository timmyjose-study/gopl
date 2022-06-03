package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Printf("%v = %v\n", tempconv.AbsoluteZeroC, tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Printf("%v = %v\n", tempconv.FreezingC, tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("%v = %v\n", tempconv.BoilingC, tempconv.CToF(tempconv.BoilingC))
}