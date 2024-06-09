package main

import "fmt"

type Address struct {
	Street string
	City   string
}

func main() {
	//function new membuat pointer yang mengacu ke data awal kosong
	a := new(Address)
	b := a

	b.City = "Bengkulu"
	b.Street = "Kepahiang"
	fmt.Println(a.City, a.Street) // Bengkulu Kepahiang
}
