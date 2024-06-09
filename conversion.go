package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var nama string = "Kerel khalif afif"
	var e = nama[0]
	fmt.Println(string(e))
}
