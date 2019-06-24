package main

import (
	"fmt"
)

func main() {
	j := 25

	if i := 22; j == 23 {
		fmt.Println("I'm in here")
		fmt.Println("Value of i", i)
	} else if i == 23 {
		fmt.Println("i was 23")
	} else {
		fmt.Println("We couldn't get any value right")
	}
	// fmt.Println("i is out of scope here", i)

	//  Missing switch expression defaults to `true`
	switch {
	case 1 == 2:
		fmt.Println(" One is equal to Two")
	case 1 != 2:
		fmt.Println("One is not equal to Two")
		fallthrough
	case 2 != 3:
		fmt.Println("Two is not equal to Three")
		fallthrough
	case 1 == 2:
		fmt.Println("One is equal to Two. THIS WILL PRINT BECAUSE OF THE fallthrough")
		fallthrough
	default:
		fmt.Println("No cases passed. THIS WILL ALSO PRINT BECAUSE OF THE fallthrough")
	}

	switch "TUX" {
	case "MR. PENGUIN", "BOND", "TUXEDO":
		fmt.Println("This is not James Bond")
	case "TUX":
		fmt.Println("This is TUX")
	}

}
