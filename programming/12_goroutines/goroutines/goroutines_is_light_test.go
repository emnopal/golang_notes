package goroutines

// demonstrating how light is goroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNum(number int) {
	fmt.Println("Display", number)
}

func TestHowLightIsGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ { // creating 1 million goroutines, 2kb*1_000_000 => near 2gb
		go DisplayNum(i) // not sequential running, since this laptop is multicore
	}
	time.Sleep(10 * time.Second)
}
