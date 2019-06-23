package main

import (
	"fmt"
)

// When you want to create a variable outside the function body,
// the short declaration operator will not work. Instead use 'var'
func main() {
	var a int = 5
	var b int = 6

	go func() {
		b = a * b
	}()

	a = b * b

	fmt.Println("Hit Enter when you want to see the answer")
	fmt.Scanln()

	fmt.Printf("a = %d, b = %d\n", a, b)
}
