package main

import (
	"fmt"
	"runtime"
)

var x int
var y string
var z bool

type flurkin int

var goose flurkin

func main() {

	// CUSTOM TYPE AND UNDERLYING TYPE
	fmt.Println("\nCUSTOM TYPE AND UNDERLYING TYPE")

	fmt.Println(goose)
	fmt.Printf("%T\n", goose)

	// CONVERSION
	fmt.Println("\nCONVERSION")

	goose = 42
	var number int
	number = int(goose)

	fmt.Printf("%d\n", number)

	// USING SPRINTF
	fmt.Println("\nUSING SPRINTF")

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	// ZERO VALUE EXAMPLE
	fmt.Println("\nZERO VALUE EXAMPLE")
	fmt.Println("x:", x, "\ny:", y, "\nz:", z)

	// OS AND ARCH
	fmt.Println("\nOS AND ARCH")

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	// BYTE AND RUNE
	// byte = uint8
	// rune = int32

}
