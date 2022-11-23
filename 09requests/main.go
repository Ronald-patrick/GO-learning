package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://ronald-patrick.tech/"

func main() {
	fmt.Println("Web requests")

	res,err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n",res)

	defer res.Body.Close() // close the connection

	databytes,err := ioutil.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}