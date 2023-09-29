package main

import "fmt"

func main() {
	// Membuat sebuah slice dengan panjang awal 5
	mySlice := make([]int, 5)

	// Menambahkan data ke dalam slice
	for i := 0; i < 3; i++ {
		newData := i + 10 // Data yang akan ditambahkan
		mySlice = append(mySlice, newData)
	}

	// Menampilkan isi slice
	fmt.Println("Slice setelah ditambahkan data:")
	fmt.Println(mySlice)
}
