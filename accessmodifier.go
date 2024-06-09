package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	//huruf besar didepan accessnya public bisa dipakai di luar package
	result := helper.Version()

	//huruf kecil didepan accessnya private tidak bisa dipakai di luar package
	// fmt.Println(helper.vesion)
	fmt.Println(result)
}
