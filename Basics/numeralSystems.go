package main

import (
	"fmt"
)

// Typed and untyped constants
const (
	a int = 43
	b     = 56
)

// IOTA
const (
	_ = 2016 + iota
	x = 2016 + iota
	y = 2016 + iota
	z = 2016 + iota
)

func main() {

	// PRINT OUT DECIMAL, BINARY AND HEX
	const integer = 42
	fmt.Printf("Decimal: %d\n", integer)
	fmt.Printf("Hex: %#X\n", integer)
	fmt.Printf("Binary: %b\n", integer)

	// BIT SHIFTING
	const shiftedInteger = 42 << 1
	fmt.Printf("\nDecimal: %d\n", shiftedInteger)
	fmt.Printf("Hex: %#X\n", shiftedInteger)
	fmt.Printf("Binary: %b\n", shiftedInteger)

	// RAW STRING LITERALS

	const literal = `
		Hello
		World
	`
	fmt.Println(literal)

	// IOTA EXAMPLE
	fmt.Println(x, y, z)
}
