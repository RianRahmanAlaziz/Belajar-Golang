package main

import "fmt"

func main() {
	name := "rian"

	switch name {
	case "rian":
		fmt.Println("hello rian")
	case "budi":
		fmt.Println("hello budi")
	default:
		fmt.Println("hello kenalan")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}
}
