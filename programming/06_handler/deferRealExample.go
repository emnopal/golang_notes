package main

import (
	"fmt"
	"time"
)

func main() {

	// set start time
	start := time.Now()

	// execute
	var sum float64 = 0
	for i := 0; i <= 1_000_000; i++ {
		sum += float64(i*10) / 18.7
	}

	defer func(start time.Time, result float64) {
		// set end time
		end := time.Since(start)
		fmt.Println("Result: ", result)
		fmt.Println("Time elapsed: ", end.Seconds(), "s")
	}(start, sum)

}
