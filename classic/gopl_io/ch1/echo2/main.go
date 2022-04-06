package echo2

import (
	"fmt"
	"os"
)

func Echo() {
	s, step := "", ""

	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}
