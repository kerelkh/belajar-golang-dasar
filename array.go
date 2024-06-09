package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "kerel"
	names[1] = "khalif"
	names[2] = "afif"

	fmt.Println(names[0], names[1], names[2])

	var nilai = [4]int{48, 40, 33, 23}
	fmt.Println(nilai)
	fmt.Println(len(nilai))

	var nilai2 = [...]int{32, 33, 23, 343, 432}
	fmt.Println(nilai2)
}
