package dup3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Dup() {
	args := os.Args[1:]
	counts := make(map[string]int)

	for _, file := range args {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Fprintf(os.Stdout, "%s:%d\n", line, n)
	}
}
