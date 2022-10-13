package main

import "fmt"

func main() {

	// Numeric casting
	var value32 int32 = 1234
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value64)

	fmt.Println(value16)
	fmt.Println(value32)
	fmt.Println(value32)

	// example of overflow
	var value16_2 int16 = 12345
	var value8_2 int8 = int8(value16_2) // this will lead to overflow

	fmt.Println(value8_2) // 57 instead of 12345 because overflow

	// byte conversion to string
	// since get value in golang from string will return byte instead of string
	// you can casting it back into string

	var test_str string = "this is string"
	get_char := test_str[3]

	fmt.Println(get_char) // 115 instead of s

	// to bring it back to string
	fmt.Println(string(get_char)) // s

}
