package main

import "fmt"

func main() {
	age := 27

	if age > 27 {
		fmt.Println("You are older than 27")
	} else if age == 27 {
		fmt.Println("You are 27 years old")
	} else {
		fmt.Println("You are younger than 27")
	}

	//short statemen
	if even := 27 % 2; even == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
