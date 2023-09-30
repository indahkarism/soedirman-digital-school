package main

import "fmt"

func main() {
	person := map[string]string{
		"flower": "jasmine",
		"tea":    "gopeks",
	}

	company := map[string]string{
		"drink": "tea",
		"brand": "gopeks",
	}
	fmt.Println(person["flower"])
	fmt.Println(person["tea"])

	fmt.Println(company["drink"])
	fmt.Println(company["brand"])
}
