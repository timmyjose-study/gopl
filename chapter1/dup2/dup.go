package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	if len(os.Args[1:]) == 0 {
		processInput(os.Stdin, counts)
	} else {
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}

			defer file.Close()
			processInput(file, counts)
		}

	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d: %s\n", count, line)
		}
	}
}

func processInput(in *os.File, m map[string]int) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		m[scanner.Text()]++
	}
}