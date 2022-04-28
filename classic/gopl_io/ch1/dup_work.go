package ch1

import (
	"bufio"
	"fmt"
	"os"
)

type fileLog struct {
	filename []string
	count    int
}

func Dup() {
	args := os.Args[1:]
	var logs map[string]fileLog
	for _, fileName := range args {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		handle(f, logs)
	}
	fmt.Println(logs)
}

func handle(f *os.File, l map[string]fileLog) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		tmp := l[input.Text()]
		tmp.count += 1
		tmp.filename = append(tmp.filename, f.Name())
		fmt.Println(tmp)
		os.Exit(222)
		l[input.Text()] = tmp
	}
}
