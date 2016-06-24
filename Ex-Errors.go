package main

import (
	"fmt"
	"math"
)
// COULD NOT figure out how to make Sqrt function -> so copied it from the URL - https://gist.github.com/zyxar/2317744 - SO that I could complete this exercise(errors).

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, string) {
	if x < 0 {
		return float64(0), ErrNegativeSqrt(x).Error()
	}
	z := float64(2.)
	s := float64(0)
	for {
		z = z - (z*z - x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}
	return s, ""
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
