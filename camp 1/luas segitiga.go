package main

import "fmt"

func main() {

	var luas, a, t float32

	fmt.Print("Masukkan panjang alas segitga: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&t)

	luas = 0.5 * a * t

	fmt.Println("Luas segitiga adalah", luas)

}
