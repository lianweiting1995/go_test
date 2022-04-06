package main

import (
	"fmt"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo2"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo3"
)

func main() {
	//ch1_func()
}

func ch1_func() {
	// classic/gopl_io/ch1/echo1
	fmt.Println("gopl_io/ch1/echo1:")
	echo1.Echo()
	// classic/gopl_io/ch1/echo2
	fmt.Println("gopl_io/ch1/echo2:")
	echo2.Echo()
	// classic/gopl_io/ch1/echo3
	fmt.Println("gopl_io/ch1/echo3:")
	echo3.Echo()
	// classic/gopl_io/ch1
	fmt.Println("gopl_io/ch1:")
	ch1.Work()
}
