package ch2

import "fmt"

type Celsius float64
type Fahrenherit float64

const (
	AbsoluteZeroC Celsius = -367.98
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenherit {
	return Fahrenherit(c*9/5 + 32)
}

func FToC(f Fahrenherit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g CC", c)
}
