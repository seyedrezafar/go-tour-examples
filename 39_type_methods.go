package main

import (
	"fmt"
	"math"
)

// the Abs method can not be directly defined on float64 because it is defined in another package 
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
