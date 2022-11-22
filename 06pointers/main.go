package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers Moduke")

	var ptr *int;
	var ptr2 *string;

	var x = 10;
	var s = "patrick";

	ptr = &x;
	ptr2 = &s;

	*ptr = 20;
	*ptr2 = "Patrick";

	fmt.Println("Value of variable is:",x)
	fmt.Println("Value of variable 2 is:",s)

}