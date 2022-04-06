package ch1

import (
	"fmt"
	"os"
)

func Work() {
	for index, arg := range os.Args[0:] {
		fmt.Println(index, arg)
	}
}
