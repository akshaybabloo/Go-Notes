package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

type Cow struct {
}

func (d Dog) Speak() string  {
	return "woof!"
}

func (c Cat) Speak() string  {
	return "meow!"
}

func (c Cow) Speak() string  {
	return "moo!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals{
		fmt.Println(animal.Speak())
	}
	
}
