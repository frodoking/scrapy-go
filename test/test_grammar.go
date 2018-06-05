package main

import "log"

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

func animal(animal Animal) {
	log.Println(animal.Speak())
}

func testInterface() {
	dog := &Dog{}
	animal(dog)
	cat := &Cat{}
	animal(cat)
}

func main() {
	testInterface()
}
