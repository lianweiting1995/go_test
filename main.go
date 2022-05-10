package main

import (
	"fmt"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo2"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo3"
	"os"
)

func main() {
	//ch1_echo()
	//ch1_dup()
	lissajous1()
}

func ch1_echo() {
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
	ch1.EchoWork()
}

func ch1_dup() {
	// classic/gopl_io/ch1/dup1
	// fmt.Println("classic/gopl_io/ch1/dup1:")
	// dup1.Dup()
	//fmt.Println("classic/gopl_io/ch1/dup2:")
	//dup2.Dup()
	//fmt.Println("classic/gopl_io/ch1/dup3:")
	//dup3.Dup()
	fmt.Println("classic/gopl_io/dup_word:")
	ch1.Dup()
}

func lissajous1() {
	//lissajous12.Lissajous(os.Stdout)
	ch1.Lissajous(os.Stdout)
}
