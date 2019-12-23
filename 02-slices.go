// https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slices := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		x := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			fx := i^j
			x[j] = uint8(fx)
		}
		slices[i] = x
	}
	return slices
}

func main() {
	pic.Show(Pic)
}
