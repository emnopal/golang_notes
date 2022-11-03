package main

// in golang, in one package, is forbidden to use same function name

/*
This is Type, Variable and Const section

Two kind of numeric type:
1. Integer ((u)int8, (u)int16, (u)int32, (u)int64)
1.5 -> uintptr (*uint) = unsigned integer. type which is large enough to hold any pointer address
2. Floating Point (float32, float64, complex64 (2x (real+imag) float32), complex128 (2x (real+imag) float64))
3. alias byte => unit8, rune => int32, int => flexible (int32 or int64), uint => flexible (uint32 or uint64)

Other Type:
Boolean, String

String method -> mostly same with python
length of string -> len(string Str): int
get character using slicing -> Str[int Number]: byte -> need to conversion from byte to string
*/

import "fmt"

// in golang, linter is strict, if import is unused, they will throw error
// import "<example of unused module>" // this is will throw error

func main() {

	// golang is strict in linter
	// if there are variables which not used, they will throw an error
	// var test string = "test" // this is will throw error, since nobody use variable test

	// _type string_
	// example of string method
	str_method := "This is string"
	fmt.Println(len(str_method)) // 14
	fmt.Println(str_method[0])   // 84 (result is in binary, need to convert into string to make human readable)
	fmt.Println(str_method[1])   // 104 (result is in binary, need to convert into string to make human readable)

}
