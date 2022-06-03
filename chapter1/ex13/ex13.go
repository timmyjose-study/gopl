package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	processArgs1()
	fmt.Printf("[processArgs1] %v\n", time.Since(start))

	start = time.Now()
	processArgs2()
	fmt.Printf("[processArgs2] %v\n", time.Since(start))
}

func processArgs1() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func processArgs2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}