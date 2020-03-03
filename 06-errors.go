package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", int(float64(e)))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	z1 := 1.0
	i := 1
	for ; i < 10; {
		z1 -= (z*z - x) / (2*z)
		
		if math.Abs(z-z1) < 10e-6 {
			fmt.Printf("Iterations: %d\n",i)
			fmt.Printf("Standard library value: %f\n", math.Sqrt(x))
			return z1, nil
		}
		z = z1
		i++
	}	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
