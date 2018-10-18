// https://tour.golang.org/moretypes/23

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Split(s, " ")
	result := make(map[string]int)

	for _, v := range arr {
		_, ok := result[v]
		if ok {
			result[v]++
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
