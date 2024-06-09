package main

import "fmt"

func main() {
	//closure
	counter := 0
	increment := func() {
		fmt.Println(counter)
		counter++
	}

	increment()
	increment()

}
