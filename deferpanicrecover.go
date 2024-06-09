package main

import "fmt"

func logging() {
	fmt.Println("Run Logging")

	message := recover()
	fmt.Println(message)
}

func sayHello() {
	defer logging()
	fmt.Println("Hello")
}

func cekNama(nama string) {
	defer logging()
	if nama == "anjing" {
		panic("Nama tidak boleh anjing")
	} else {
		fmt.Println("Hai,", nama)
	}
}

func main() {
	sayHello()
	cekNama("anjing")
}
