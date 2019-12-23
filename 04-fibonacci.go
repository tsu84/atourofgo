// https://tour.golang.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib0 := 0
	fib1 := 0
	return func() int {
		if fib1 == 0 {
			fib1 = 1
			return 1
		}
		fib := fib1 + fib0
		fib0 = fib1
		fib1 = fib
		
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
