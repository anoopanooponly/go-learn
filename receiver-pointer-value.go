package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {

	// The rule about pointers vs. values for receivers is that value methods 
	// can be invoked on pointers and values, but pointer methods can only be invoked on pointers
	
	// new(Dog) is equal to &Dog{} returns a pointer
	// here new(Dog) can invoke speak of Dog with or without pointer in speak
	// if we replace new(Cat) with Cat{} it will throw error
	animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
	// animals := []Animal{new(Dog), Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
