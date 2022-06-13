package ch4

import (
	"fmt"
	"testing"
)

func TestRev(t *testing.T) {
	var s = []int{1, 4, 5, 7, 9, 100}
	reverse(s)
}

func TestRev1(t *testing.T) {
	s := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s[:])
	fmt.Println(s)
}
