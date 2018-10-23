package main

import (
	"fmt"
	"net/http"
)

type Hello struct {}

// The name of the function should be this.
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:1313", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
