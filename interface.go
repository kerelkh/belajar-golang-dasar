package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
	Age  int
}

func (person Person) GetName() string {
	return person.Name
}

func sayHello(value HasName) {
	fmt.Println("Hai,", value.GetName())
}

func main() {
	person := Person{"John", 30}
	sayHello(person)
}
