package go_channel

import (
	"fmt"
	"testing"
)

func FibonacciMain(c, quit chan int) {
	fmt.Println("fibonacci started")
	x, y := 0, 1 // assign x and y to 0 and 1
	for {        // while true
		fmt.Printf("x = %d\n", x)
		fmt.Printf("y = %d\n", y)
		select { // switch case
		case c <- x: // assign x into c
			x, y = y, x+y // update x and y value
		case quitValue := <-quit: // assign quit channel into quitValue
			fmt.Printf("quitValue = %d\n", quitValue)
			return // finish
		}
	}
}

func TestFibonacciMain(t *testing.T) {
	c := make(chan int)
	q := make(chan int)
	go func() {
		fmt.Println("Goroutine started")
		for i := 0; i < 5; i++ {
			value := <-c
			fmt.Printf("received %d\n", value)
		}
		q <- 999
	}()
	fmt.Println("calling FibonacciMain()")
	FibonacciMain(c, q)
}
