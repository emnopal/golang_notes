package main

import "fmt"

func main() {

	// _variable_
	// variable (in golang variable is static, you cannot combine another data type into another in variable)
	var example_vars string = "This is variable"
	fmt.Println(example_vars)

	// above is best practice if you want to write variable with explicit data type

	// or you can split it into
	var example_vars2 string
	example_vars2 = "This is variable"
	fmt.Println(example_vars2)

	// replace (not reinitialize variable), remember: the type must be same (string)
	// variable is unique, you can't reinitialize the variable
	example_vars = "This is second variable" // = 21 // will throw an error

	fmt.Println(example_vars)

	// another way to write variable in golang
	// without explicitly write the data types

	var second_example_vars = "This is string"
	fmt.Println(second_example_vars)

	// for initialize variable like this
	// the statement must be one line

	// error
	// var test
	// test = 12
	// fmt.Println(12) // this is will throw error
	// instead of that you must specify data type:

	// not error
	var test int
	test = 12
	fmt.Println(test) // ok

	// another way to write variable,
	// this is best practice if developer don't want to write data type
	third_example_vars := "This is third example"
	fmt.Println(third_example_vars)

	// this is will throw error, for vars with := you don't need to write data type
	// forth_example_vars string := "This is forth example"

	// writing multiple variable, this is best practice for writing multiple variable
	var (
		test0       = "test0"
		test1       = "test1"
		test2       = "test2"
		test3 uint8 = 12 // can use different data type
	)

	fmt.Println(test0)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)

}
