package dup2

import (
	"bufio"
	"fmt"
	"os"
)

func Dup() {
	args := os.Args[1:]
	counts := make(map[string]int)

	if len(args) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLine(f, counts)
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLine(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		count[input.Text()]++
	}
}
