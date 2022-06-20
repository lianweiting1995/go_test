package ch4

import (
	"bufio"
	"fmt"
	"os"
)

func Scanword() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		fmt.Fprintf(os.Stdout, "out: %s \n", in.Text())
	}
}
