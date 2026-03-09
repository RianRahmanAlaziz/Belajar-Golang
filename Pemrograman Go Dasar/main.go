package main

import "fmt"

// variable

// func main() {
// 	var firstName string = "jhon"

// 	var lastName string
// 	lastName = "wick"

// 	fmt.Printf("halo %s %s!\n", firstName, lastName)
// }

// Type Data
// Tipe Data Numerik Non-Desimal

// func main() {
// 	var positiveNumber uint8 = 89
// 	var negativeNumber = -1243423644

// 	fmt.Printf("bilangan positif: %d\n", positiveNumber)
// 	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

// }

// Tipe Data Numerik Desimal
// func main() {
// 	var decimalNumber = 2.62

// 	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
// 	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

// }

// Tipe Data bool (Boolean)

// func main() {
// 	var exist bool = true
// 	fmt.Printf("exist? %t \n", exist)

// }

// Tipe Data string

// func main() {
// 	var message = `Nama saya "Rian".
// Salam kenal.
// Mari belajar "Golang".`

// 	fmt.Println(message)

// }

// Penggunaan Konstanta

// func main() {
// 	const firstName = "john"
// 	fmt.Print("halo ", firstName, "!\n")

// 	fmt.Println("john wick")
// 	fmt.Println("john", "wick")

// 	fmt.Print("john wick\n")
// 	fmt.Print("john ", "wick\n")
// 	fmt.Print("john", " ", "wick\n")

// }

// Deklarasi Multi Konstanta

// func main() {
// 	const (
// 		square         = "kotak"
// 		isToday  bool  = true
// 		numeric  uint8 = 1
// 		floatNum       = 2.2
// 	)

// 	/*
// 		square, dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "kotak"
// 		isToday, dideklarasikan dengan metode manifest typing dengan tipe data bool dan nilainya true
// 		numeric, dideklarasikan dengan metode manifest typing dengan tipe data uint8 dan nilainya 1
// 		floatNum, dideklarasikan dengan metode type inference dengan tipe data float dan nilainya 2.2
// 	*/

// 	const (
// 		today string = "senin"
// 		sekarang
// 		isToday2 = true
// 	)

// 	/*
// 		today dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
// 		sekarang dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
// 		isToday2 dideklarasikan dengan metode type inference dengan tipe data bool dan nilainya true
// 	*/

// 	const satu, dua = 1, 2
// 	const three, four string = "tiga", "empat"

// 	/*
// 		satu, dideklarasikan dengan metode type inference dengan tipe data int dan nilainya 1
// 		dua, dideklarasikan dengan metode type inference dengan tipe data int dan nilainya 2
// 		three, dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "tiga"
// 		four, dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "empat"
// 	*/

// }

func main() {
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}

}
