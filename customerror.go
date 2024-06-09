package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (nf *notFoundError) Error() string {
	return nf.Message
}

func saveData(s string) error {
	if s == "" {
		return &validationError{Message: "Data harus diisi"}
	}

	if s != "kerel" {
		return &notFoundError{Message: "Data tidak ditemukan"}
	}

	return nil
}

func main() {
	err := saveData("kerel")
	if err != nil {
		switch e := err.(type) {
		case *validationError:
			fmt.Println("Error Validation:", e.Message)
		case *notFoundError:
			fmt.Println("Error Not Found:", e.Message)
		default:
			fmt.Println("Error lain:", err)
		}
	} else {
		fmt.Println("Data berhasil disimpan")
	} 
}
