package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello")
}

// function parameter

func sayHelloTo(fristName string, Lastname string) {
	fmt.Println("hello", fristName, Lastname)
}

// function return value

func getHello(name string) string {
	hello := "hello " + name
	return hello
}

// Returning Multiple Values

func getFullName() (string, string) {
	return "rian", "rahman"
}

// named Return Values

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "rian"
	middleName = "rahman"
	lastName = "al aziz"

	return firstName, middleName, lastName
}

// variadic function

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

// function as value

func getGoodBye(name string) string {
	return "Good Bye" + name
}

// function as Parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// Anonymous Function

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Blacklist ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// Recursive Function

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(10))
}
