package main

import "fmt"

// Slices are built over arrays.

func main() {

	//  Create slice using Composite Literals
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice, len(slice))
	fmt.Printf("Type: %T\n", slice)

	fmt.Println()
	// Creating using make
	makeSlice := make([]int, 10, 12)
	fmt.Println("Before appending:", makeSlice, len(makeSlice), cap(makeSlice))
	makeSlice = append(makeSlice, 11, 12, 13)
	fmt.Println("After appending:", makeSlice, len(makeSlice), cap(makeSlice))
	// The goruntime automatically increases the capacity by 2 by copying over all the elements
	// from the first array to the second when the array length exceeds the capacity

	fmt.Println()
	// Acessing slice using for..range
	for index, value := range slice {
		fmt.Println(index, value)
	}

	fmt.Println()
	// Slicing a slice
	fmt.Println(slice[1:3])

	fmt.Println()
	// Appending a slice
	newSlice := append(slice, 6, 7, 8)
	fmt.Println(newSlice)

	fmt.Println()
	// Variadic parameter
	secondSlice := append(slice[2:3], newSlice...)
	fmt.Println(secondSlice)

	fmt.Println()
	// Deleting from a slice
	secondSlice = append(secondSlice[:2], secondSlice[5:]...)
	fmt.Println(secondSlice)

}
