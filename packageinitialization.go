package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
	fmt.Println("My connection is :" + database.GetConnection())
}
