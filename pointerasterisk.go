package main

import "fmt"

type Address struct {
	Street string
	City   string
}

func main() {
	//jika variabel pointer dipindahkan maka yang berubah hanya variable itu saja tapi jika ingin membuat
	//semua variabel pointer yang mengacu ke pointer tersebut berubah maka menggunakan asterisk (*)
	a := Address{"Jalan", "Jakarta"}
	b := &a
	c := b
	fmt.Println(a, b, c)
	a.Street = "Jalan Baru"
	fmt.Println(a, b, c)

	*b = Address{"Kuterejo", "Bengkulu"}
	fmt.Println(a, b, c)

}
