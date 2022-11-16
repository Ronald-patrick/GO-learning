package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Enter your name ")
	var reader = bufio.NewReader(os.Stdin);

	// comma ok || error ok

	var input,_ = reader.ReadString('\n')
	fmt.Println("Welcome to GOlang : ",input) // Type string
}