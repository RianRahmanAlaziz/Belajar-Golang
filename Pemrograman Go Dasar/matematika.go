package main

import "fmt"

func main() {

	var a = 10
	var b = 20
	var d = 15
	var c = a + b - d

	fmt.Println(c)

	// Augmented Assignments
	var i = 10

	i += 10 // i = i + 10
	fmt.Println(i)
	i += 5 // i = i + 5
	fmt.Println(i)

	// Unary Operator

	var j = 1

	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
}
