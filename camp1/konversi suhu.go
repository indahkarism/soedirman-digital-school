package main

import (
	"fmt"
)

func main() {
	suhu()
}
func suhu() {
	var (
		c, r, f, k float64
	)
	fmt.Print("Masukkan suhu celcius = ")
	fmt.Scanln(&c)
	reamur := (c * 4) / 5
	fmt.Println(c, "C =", reamur, "R")
	kelvin := c + 273
	fmt.Println(c, "C =", kelvin, "K")
	fah := (c*9)/5 + 32
	fmt.Println(c, "C =", fah, "F")

	fmt.Print("Masukkan suhu Reamur = ")
	fmt.Scanln(&r)
	celcius := (r * 5) / 4
	fmt.Println(r, "R =", celcius, "C")
	kelvin1 := (r/4)/5 + 273.15
	fmt.Println(r, "R =", kelvin1, "K")
	fah1 := (r*9)/4 + 32
	fmt.Println(r, "R =", fah1, "F")

	fmt.Print("Masukkan suhu fahrenheit = ")
	fmt.Scanln(&f)
	celcius1 := (f - 32) * 1.8
	fmt.Println(f, "F =", celcius1, "C")
	kelvin2 := (f - 459.67) / 1.8
	fmt.Println(f, "F =", kelvin2, "K")
	reamur1 := (f - 32) / 2.25
	fmt.Println(f, "F =", reamur1, "R")

	fmt.Print("Masukkan suhu Kelvin = ")
	fmt.Scanln(&k)
	celcius2 := (k - 273.5) * 1.8
	fmt.Println(k, "K =", celcius2, "C")
	reamur2 := (k - 273.15) * 0.8

	fmt.Println(k, "K =", reamur2, "R")
	fah2 := (k * 1.8) - 459.67
	fmt.Println(k, "K =", fah2, "F")

}
