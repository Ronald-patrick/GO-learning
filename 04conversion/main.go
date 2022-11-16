package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Conversions")

	var reader = bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	i1,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println(i1*2)
	}

}
