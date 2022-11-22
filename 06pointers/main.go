package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int;

	var x = 10;

	ptr = &x;

	*ptr = 20;

	fmt.Println("Value of variable is ",x)

}