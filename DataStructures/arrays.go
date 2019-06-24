package main

import "fmt"

// Use slices instead of array in GoLang
func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i * 5
	}

	fmt.Println(arr)
	fmt.Printf("\nType of arr: %T\n", arr)

	// Length of the array is part of the type :O
	var arr1 [6]int
	fmt.Printf("\nType of arr1: %T\n", arr1)
}
