package main

import (
	"fmt"
	"os"
)

func main() {
	// ambil data argument
	fmt.Println("os Arguments: ")
	args := os.Args
	fmt.Println(args)
	fmt.Println()

	// ambil nama host
	fmt.Println("os Hostname: ")
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println()

	// read more: https://pkg.go.dev/os
}
