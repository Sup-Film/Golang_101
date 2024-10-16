package main

import (
	"fmt"
)

// Define the Speaker interface
type Speaker interface {
	Speak() string
}
 
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Person struct {
	Name string
}

func (p Person) Speak() string  {
	return "Hello!"
}

func makeSound(s Speaker)  {
	fmt.Println(s.Speak())
}

func main()  {
	dog := Dog{Name: "Dog"}
	person := Person{Name: "Person"}

	makeSound(dog)
	makeSound(person)
}