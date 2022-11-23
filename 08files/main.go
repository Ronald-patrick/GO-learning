package main

import (
	"fmt"
	"time"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to files in golang")

	filename := "./log.txt"

	file,err := os.OpenFile(filename,os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	content := "LoggedIn at: " + time.Now().Format("01-02-2006 15:04:05 Monday") + " \n"
	_, err = io.WriteString(file,content)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// readFile(filename)

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(databyte)

	fmt.Println("Log Data :\n"+ strings.TrimSpace(data))
}