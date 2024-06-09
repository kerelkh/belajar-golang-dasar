package main

import "fmt"

func main() {
	for counter := 1; counter < 10; counter++ {
		fmt.Println("Counter is", counter)
	}

	names := []string{"kerel", "khalif", "afif"}

	for _, name := range names {
		fmt.Println("Name is", name)
	}
}
