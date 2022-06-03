package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		contents, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		for _, line := range strings.Split(string(contents), "\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d: %v\n", count, line)
		}
	}
}