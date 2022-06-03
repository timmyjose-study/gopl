package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d: %s\n", count, line)
		}
	}
}