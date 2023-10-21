package main

import "fmt"

func main() {

	var r, luas, phi float32

	phi = 3.14

	fmt.Print("Masukkan panjang jari-jari lingkaran: ")
	fmt.Scanln(&r)

	luas = phi * r * r

	fmt.Println("Luas lingkaran adalah ", luas)

}
