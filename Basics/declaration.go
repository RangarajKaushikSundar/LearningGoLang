package main

import "fmt"

var a, b int

func main() {
	a = 1
	b = -1

	fmt.Println("a: ", a, ", b:", b)
	// Swap values
	a, b = b, a
	fmt.Println("a: ", a, ", b:", b)

}
