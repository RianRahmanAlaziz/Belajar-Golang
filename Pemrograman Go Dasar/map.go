package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["nama"] = "rian"
	// person["address"] = "jauh"

	person := map[string]string{
		"name":    "rian",
		"address": "jauh",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "rian"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
