// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z1 := 1.0
	i := 1
	for ; i < 10; {
		z1 -= (z*z - x) / (2*z)
		
		if math.Abs(z-z1) < 10e-6 {
			fmt.Printf("Iterations: %d\n",i)
			fmt.Printf("Standard library value: %f\n", math.Sqrt(x))
			return z1
		}
		z = z1
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
