package main

import "fmt"

func main() {
	var array = [...]string{"kerel", "khalif", "afif", "joko", "budi", "nugraha"}
	var slice = array[4:6]

	fmt.Println(slice)

	slice2 := array[:3]
	fmt.Println(slice2)

	slice3 := array[2:]
	fmt.Println(slice3)

	slice4 := array[:]
	fmt.Println(slice4)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//append membuat slice baru dengan menambahkan data diposisi akhir, jika penuh maka dibuat array baru
	var days = [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	var daysSlice = days[4:7]
	fmt.Println(daysSlice)
	daysSlice[0] = "sholat jumat"
	fmt.Println(daysSlice)
	fmt.Println(days)

	daysSlice2 := append(daysSlice, "Haloo")
	fmt.Println(daysSlice2)
	fmt.Println(days)

	//buat slide dari array baru
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "test"
	newSlice[1] = "tist"

	fmt.Println(newSlice)

	//ini deklarasi array
	var newArray = [...]int{4, 3, 2, 3}

	//ini deklarasi slice
	var newSlices = []int{3, 2, 3, 4}
	fmt.Println(newArray)
	fmt.Println(newSlices)
}
