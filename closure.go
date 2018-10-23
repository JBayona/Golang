// https://tour.golang.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	fib := 0
	start := 1
	end := 1
	return func() int {
		if count == 0 {
			fib = 0
		} else if count <= 2 {
			fib = 1
		} else {
			fib = start + end
			end = start
			start = fib
		}
		count++
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
