package main

import "fmt"

func main() {
	// first: assign input0, 1 and 2 to string
	var (
		input0, input1, input2 string
	)

	// by default, it will be empty string
	fmt.Println(input0) // ''
	fmt.Println(input1) // ''
	fmt.Println(input2) // ''

	// this, will change the input0,1 and 2 pointer from '' to user input
	fmt.Println("Input 1:")
	_, err0 := fmt.Scanln(&input0)

	fmt.Println("Input 1:")

	_, err1 := fmt.Scanln(&input1)

	fmt.Println("Input 1:")
	_, err2 := fmt.Scanln(&input2)

	// error handling
	if err0 != nil {
		fmt.Println(err0)
	}
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	// changed input0, 1 and 2 pointer
	fmt.Println(input0)
	fmt.Println(input1)
	fmt.Println(input2)
}
