package main

import (
	"fmt"
	"os"

	"github.com/lianweiting1995/go_test/ch3"
	"github.com/lianweiting1995/go_test/ch4"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo2"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/echo3"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/fetch1"
	"github.com/lianweiting1995/go_test/classic/gopl_io/ch1/server3"
)

func main() {
	//ch1_echo()
	//ch1_dup()
	//lissajous1()
	//fetch()
	// server()
	// mandelbrotServer()
	// dedup()
	// charcount()
	// scanword()
	// issue()
	//issueAndTemplate()
	findLink()
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

func fetch() {
	fetch1.Fetch()
	//ch1.FetchWork()
	// fetch_all1.FetchAll1()
}

func server() {
	// server1.Server()
	// server2.Server()
	server3.Server()
}

func mandelbrotServer() {
	ch3.HttpServer()
}

func dedup() {
	ch4.Dedup()
}

func charcount() {
	ch4.CharCount()
}

func scanword() {
	ch4.Scanword()
}

func issue() {
	ch4.IssueFunc()
}

func issueAndTemplate() {
	ch4.InitTemplate()
}

func findLink() {
	ch4.FindLink()
}
