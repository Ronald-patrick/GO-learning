package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Module")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println(time.Date(2020, time.December, 11, 22, 23, 0, 0, time.Local))
}

// Build with go build GOOS="windows" go build
