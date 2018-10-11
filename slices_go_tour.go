package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	xs := make([]([]uint8), dy)
	for i := range xs {
		xs[i] = make([]uint8, dx)
		for j := range xs[i] {
			xs[i][j] = uint8((i & j) & (i & j))
		}
	}
	return xs
}

func main() {
	pic.Show(Pic)
}
