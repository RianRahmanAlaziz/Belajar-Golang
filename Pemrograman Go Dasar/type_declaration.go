package main

import "fmt"

func main() {

	type NoKTP string

	var ktpRian NoKTP = "11111111111"

	var contoh string = "22222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRian)
	fmt.Println(contohKtp)
}
