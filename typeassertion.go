package main

import "fmt"

func typeAssertion(name string) interface{} {
	return "OK"
}

func main() {
	result := typeAssertion("hello")
	changeType := result.(string)
	fmt.Println(changeType)

	// changeInt := result.(int)
	// fmt.Println(changeInt) // panic: interface conversion: interface {} is string, not int

	switch value := result.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("int:", value)
	default:
		fmt.Println("unknown type")

	}

}
