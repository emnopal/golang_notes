package goroutines

// for unit test: file_test.go -> add _test in every file, so golang will recognize the file as unit test file

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// running this using unit test
func TestCreateGoroutines(t *testing.T) {
	// RunHelloWorld() // without goroutine
	go RunHelloWorld() // with goroutine
	fmt.Println("Test")

	// to make this function is await, make this function sleep
	// await in here to demonstrating how asynchronous is working
	time.Sleep(1 * time.Second) // wait for 1 second (by default, golang use nanosecond)
}
