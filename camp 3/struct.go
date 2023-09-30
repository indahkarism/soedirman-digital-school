package main

import "fmt"

type costumer struct {
	name, address string
	age           int
}

func main() {
	var yana costumer
	yana.name = "yana jelek"
	yana.address = "tongsampah"
	yana.age = 1000

	fmt.Println(yana.name)
	fmt.Println(yana.address)
	fmt.Println(yana.age)
}
