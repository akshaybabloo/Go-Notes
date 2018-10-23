package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type: %T", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
