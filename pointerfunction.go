package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

func changeToIndonesia(address *Address) {
	address.State = "Indonesia"
}

func main() {
	address := Address{"kebon bawang", "Jakut", "London"}
	changeToIndonesia(&address)

	fmt.Println(address.State)
}
