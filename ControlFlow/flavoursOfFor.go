package main

import "fmt"

// FLAVOURS OF FOR

func main() {
	var i int
	// for loop
	fmt.Println("\nfor loop")
	for i = 0; i < 5; i++ {
		fmt.Printf("%d\t", i)
	}

	// while loop
	fmt.Println("\n\nwhile loop")
	for i < 10 {
		fmt.Printf("%d\t", i)
		i++
	}

	// while(true) loop
	fmt.Println("\n\nwhile(true) loop")
	for {
		if i > 14 {
			break
		} else {
			fmt.Printf("%d\t", i)
			i++
			continue
		}
	}
	fmt.Println("\n")

}
