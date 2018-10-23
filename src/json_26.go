package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	Id int
	Title string
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	jsonContent := todosJson(content)

	fmt.Println(jsonContent)
}

func todosJson(content string) []Todo {
	todos := make([]Todo, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var todo Todo

	for decoder.More() {
		err := decoder.Decode(&todo)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}
	return todos
}
