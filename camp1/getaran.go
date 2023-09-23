package main

import (
	"fmt"
)

func main() {
	frekuensi()
}

func frekuensi() {
	var (
		n, t float64
	)

	fmt.Print("Masukkan jumlah getaran: ")
	fmt.Scanln(&n)
	fmt.Print("Masukkan waktu dalam satuan sekon: ")
	fmt.Scanln(&t)
	frekuensi := n / t
	fmt.Println("Frekuensi = ", frekuensi, "hz")
}
