package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Print(os.Args[0] + ":")

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}