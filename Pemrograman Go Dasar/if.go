package main

import "fmt"

func main() {
	name := "asd"

	if name == "rian" {
		fmt.Println("hello rian")
	} else if name == "joko" {
		fmt.Println("hello Joko")
	} else {
		fmt.Println("boleh kenalan")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama benar")
	}
}
