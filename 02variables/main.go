package main

import (
	"fmt"
)

const LoginToken = "srfawedf" // This is now a public variable

func main() {
	var username string = "ronald"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallVal int = 123456
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float32 = 255.5552414124
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	// aliases

	var anotherVariable int; // 0 default
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable)

	// implicit type

	var anyType = "this is string"
	fmt.Println(anyType)

	// Without var : cannot be used in global

	numberOfUser := 30000.0 // Type decided but cannot change type
	fmt.Println(numberOfUser)




}

