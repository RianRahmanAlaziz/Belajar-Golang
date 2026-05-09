package golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGOMAXPROCS(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()

	fmt.Printf("Total CPU: %d\n", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Printf("Total Thread: %d\n", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Printf("Total Goroutine: %d\n", totalGoroutine)
}
