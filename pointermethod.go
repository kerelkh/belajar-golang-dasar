package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person *Person) changeName() {
	person.Name = "Mr." + person.Name
}

func main() {
	// make pointer method example
	person := Person{"John", 30}
	fmt.Println("Before:", person.Name)
	person.changeName()
	fmt.Println("After:", person.Name)
}
