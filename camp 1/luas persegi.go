package main

import "fmt"

func main() {

	var luas, a float32

	fmt.Print("Masukkan panjang sisi: ")
	fmt.Scanln(&a)

	luas = a * a

	fmt.Println("Luas persegi adalah", luas)

}
