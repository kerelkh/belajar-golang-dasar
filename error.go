package main

import (
	"errors"
	"fmt"
)

func pembagian(pembagi int, dibagi int) (int, error) {
	if dibagi == 0 {
		return 0, errors.New("Pembagi tidak boleh Nol")
	}

	return pembagi / dibagi, nil
}

func main() {
	pembagi := 10
	dibagi := 0
	if value, err := pembagian(pembagi, dibagi); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hasil Pembagian : ", value)
	}
}
