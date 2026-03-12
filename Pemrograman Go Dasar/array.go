package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "rian"
	names[1] = "rahman"
	names[2] = "al aziz"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)

	// function array

	var value = [...]int{
		90,
		80,
		95,
	}

	fmt.Println(value)
	fmt.Println(len(value))
	value[0] = 100
	fmt.Println(value)
}
