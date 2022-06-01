package ch2

import (
	"fmt"
	"testing"
)

func TestTempconv0(t *testing.T) {
	var c Celsius
	var f Fahrenherit
	fmt.Println(c == 0)
	fmt.Println(f == 0)
	// fmt.Println(c == f)
	fmt.Println(c == Celsius(f))
}

func TestTempconv1(t *testing.T) {
	var c Celsius = 100
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
