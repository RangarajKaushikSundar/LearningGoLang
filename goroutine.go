package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 6

	go func() {
		b = a * b
	}()

	a = b * b

	fmt.Println("Hit Enter when you want to see the answer")
	fmt.Scanln()

	fmt.Printf("a = %d, b = %d\n", a, b)
}
