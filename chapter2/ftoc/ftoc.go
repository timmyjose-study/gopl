package main

import "fmt"

const (
	freezingF = 32.0
	boilingF  = 212.0
)

func main() {
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}