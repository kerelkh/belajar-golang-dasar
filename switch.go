package main

import "fmt"

func main() {
	name := "kerel"

	switch name {
	case "kerel":
		fmt.Println("hello kerel")
	case "khliaf":
		fmt.Println("hello khliaf")
	default:
		fmt.Println("hello stranger")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("your name is too long")
	case length < 5:
		fmt.Println("your name is too short")
	default:
		fmt.Println("your name is just right")
	}
}
