package ch4

import (
	"fmt"
	"testing"
)

func TestEquoteSlice(t *testing.T) {
	var m = make(map[string]int)
	slice1 := []int{1, 2, 3, 5}
	slice2 := []string{"1", "2", "3", "4"}
	str1 := fmt.Sprintf("%q\n", slice1)
	str2 := fmt.Sprintf("%q\n", slice2)
	m[str1] = 1
	m[str2] = 2
	fmt.Println(str1, str2)
}
