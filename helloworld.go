package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界");
	fmt.Println(foo());

	for i:=0; i<10; i++ {
		if i%2 == 0 {
			fmt.Println(i);
		}
	}
}

func foo() string {
	fmt.Println("I'm in foo");
	var message = "hello world";
	return message;
}