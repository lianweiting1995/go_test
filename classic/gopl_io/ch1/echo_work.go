package ch1

import (
	"fmt"
	"os"
)

func EchoWork() {
	for index, arg := range os.Args[0:] {
		fmt.Println(index, arg)
	}
}
