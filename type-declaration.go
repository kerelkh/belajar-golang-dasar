package main

import "fmt"

func main() {
	type NoKTP string

	var noktp NoKTP = "123123123123123"
	fmt.Println(noktp)
}
