package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z :=1.0
	oldz := z
	fmt.Println(z, z*z, x-z*z)
	
	for {
		oldz = z
		z -= (z*z -x)/(2*z)
		if oldz == z {
			break
		}
		fmt.Println(z, z*z, x-z*z)

		if math.Abs(x - z*z) < 0.000000000000001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(81))
	fmt.Println(math.Sqrt(81))
	fmt.Println(Sqrt(89))
	fmt.Println(math.Sqrt(89))
}

