package main

/* package name convention see here: https://go.dev/blog/package-names */

import "fmt"

func main() {
	// golang doesn't need to add ';' after statement
	fmt.Println("Hello, World!")
}

// to compile: go build name_file.go
// to run without compile (just for development): go run name_file.go
