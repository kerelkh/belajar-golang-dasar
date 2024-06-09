package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

func main() {
	// POINTER => kemampuan membuat pass by value menjadi pass by reference jadi hemat memori (variabel)
	address := Address{"Jalan", "Jakarta", "DKI"}
	address1 := &address

	fmt.Println(address)
	fmt.Println(address1)

	address1.City = "Bengkulu"
	fmt.Println(address)
	fmt.Println(address1)

}
