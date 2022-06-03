package main

import "fmt"

const boilingF = 212.0

func main() {
	f := boilingF
	c := (f - 32.0) * 5.0 / 9.0
	fmt.Printf("The boiling point of water is %gC\n", c)
}