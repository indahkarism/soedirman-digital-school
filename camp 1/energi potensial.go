package main

import (
	"fmt"
)

func main() {
	energipotensial()
}

func energipotensial() {

	var (
		m, g, h int
	)
	fmt.Print("Masukkan massa dalam satuan kg : ")
	fmt.Scanln(&m)
	fmt.Print("Masukkan percepatan gravitasi dalam satuan m/s2: ")
	fmt.Scanln(&g)
	fmt.Print("Masukkan ketinggian benda dari permukaan tanah dalam satuan meter: ")
	fmt.Scanln(&h)
	ep := m * g * h
	fmt.Println("Energi Potensial = ", ep, "joule")
}
