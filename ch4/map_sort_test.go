package ch4

import (
	"fmt"
	"testing"
)

func TestMapSort(t *testing.T) {
	data := [...]string{
		"d",
		"e",
		"a",
		"lia",
		"ccc",
	}
	mapSort(data[:])
	fmt.Println(data)

}

func TestMap1(t *testing.T) {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	age, ok := ages["bo"]
	fmt.Println(age, ok)
}
