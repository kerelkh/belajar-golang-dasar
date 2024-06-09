package main

import "fmt"

type sayType func(string) string

func alias(name string) string {
	return "alias"
}

func sayHelloWithFuncParamAlias(name string, alias sayType) {
	fmt.Println("Hello ", alias("Kerel"))
}

func sayHello() {
	fmt.Println("Hello!")
}

func sayHelloWithParam(firstName string, lastName string) {
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func sayHelloWithReturn() string {
	return "Hello!"
}

func sayHelloWithMultipleReturn() (string, string) {
	return "Hello", "world"
}

func sayHelloMultipleIgnoreReturn() (string, string) {
	return "Hello", "world"
}

func sayHelloWithNamedReturn() (first, second string) {
	first = "Hello"
	second = "world"
	return
}

func dataNumbers(numbers ...int) {
	for _, data := range numbers {
		fmt.Println(data)
	}
}

// function as variabel
func getHello(name string) string {
	return "Hello, " + name
}

// function parameter
func sayHelloWithFuncParam(name string, filter func(string) string) {
	fmt.Println("Hai, ", filter(name))
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// function factorial recursive
func factorialRecursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorialRecursive(n-1)
	}
}

func main() {
	//function
	sayHello()

	//function with param
	sayHelloWithParam("John", "Doe")

	//return function
	fmt.Println(sayHelloWithReturn())

	//multiple return
	first, second := sayHelloWithMultipleReturn()
	fmt.Println(first, second)

	//ignore return
	_, nodata := sayHelloMultipleIgnoreReturn()
	fmt.Println(nodata)

	//named return
	first, second = sayHelloWithNamedReturn()
	fmt.Println(first, second)

	//variadic function
	dataNumbers(1, 2, 3, 4, 5)

	//send slice in variadic function
	slice := []int{10, 42, 23, 22}
	dataNumbers(slice...)

	//function as value
	hello := getHello
	fmt.Println(hello("John"))

	//function as param
	sayHelloWithFuncParam("Anjing", spamFilter)

	//or function as param with variable
	filter := spamFilter
	sayHelloWithFuncParam("Anjing", filter)

	//function with alias
	sayHelloWithFuncParamAlias("Kerel", alias)

	//use fatorial recursive function
	fmt.Println(factorialRecursive(5))

}
