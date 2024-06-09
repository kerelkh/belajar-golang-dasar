package main

import "fmt"

func sayHello() map[string]string {
	return nil
}

func main() {

	//nil hanya untuk interface, function, map, slice, array, pointer dan channel
	hasil := sayHello()

	if hasil == nil {
		fmt.Println("hasil adalah nil")
	} else {
		fmt.Println("hasil bukan nil")
	}
}
