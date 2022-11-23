package main

import (
	"fmt"
	"sort"
)

func main() {

	var numList = [4]int{1, 2, 3, 4} // Inititalization

	for i := 0; i < 4; i++ {
		numList[i] = i * 10
	}

	fmt.Println(numList)
	fmt.Println(len(numList))

	// Slices : flexible array

	fruitList := []string{"Apple", "Peach"}

	fruitList = append(fruitList, "Mango")

	fmt.Printf("Type : %T \n", fruitList)
	fmt.Println(fruitList)

	// Dynamic allocation

	highScores := make([]int, 4)
	highScores[0] = 10
	highScores[1] = 33

	sort.Ints(highScores)

	fmt.Println(highScores)

	// Remove element from slice

	index := 1
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)

	// Maps

	m := make(map[string]int)
	m["JS"] = 1
	m["C++"] = 20

	// delete(m,"JS")

	for key, value := range m {
		fmt.Println(key, value)
	}

	user1 := User{"Ronald", "ronald@gmail.com", true, 21}
	fmt.Println(user1)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
