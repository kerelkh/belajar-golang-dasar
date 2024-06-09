package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Kerel Khalif Afif",
		"age":  "27",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])

	books := make(map[string]string)
	books["title"] = "The Great Gatsby"
	books["author"] = "F. Scott Fitzgerald"
	fmt.Println(books)

	delete(books, "author")
	fmt.Println(books)
}
