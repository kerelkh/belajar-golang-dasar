package main

import "fmt"

type Customer struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}

func (person Person) sayHello() {
	fmt.Println("Hello", person.Name)
}

func (person Person) addAge(number int) int {
	umur := person.Age + number
	return umur
}

func main() {
	customer := Customer{Name: "John", Age: 30}
	fmt.Println(customer.Name) // Output: John

	var customer2 Customer
	customer2.Name = "Kerel"
	customer2.Age = 27
	fmt.Println(customer2)

	person := Person{"Kerel Khalif Afif", 27}
	person.sayHello() // Output: Hello Kerel Khalif Afif
	fmt.Println(person.addAge(3))
}
