package main

import "fmt"

func main() {
	hash := make(map[string]float32)

	hash["Jorge"] = 100
	hash["Ivan"] = 100
	hash["Maria"] = 99.9

	fmt.Println(hash)

	delete(hash, "Maria")

	for k, v := range hash {
		fmt.Println(k, ":", v)
	}
}
