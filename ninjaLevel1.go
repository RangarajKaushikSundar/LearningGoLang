package main

import "fmt"

var x int
var y string
var z bool

type flurkin int

var goose flurkin

func main() {

	// CUSTOM TYPE AND UNDERLYING TYPE
	fmt.Println(goose)
	fmt.Printf("%T\n", goose)

	// CONVERSION
	goose = 42
	var number int
	number = int(goose)

	fmt.Printf("%d\n", number)

	// USING SPRINTF
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	// ZERO VALUE EXAMPLE
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
