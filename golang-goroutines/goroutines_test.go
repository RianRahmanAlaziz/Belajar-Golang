package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWord() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWord()
	fmt.Println("Run Hello World")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
