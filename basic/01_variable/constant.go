package main

func main() {

	// _constant_
	// first time when you declared constant, you must specify the data in one line
	// and also constant is immutable, you can't change the constant data or value
	// constant if is not used, the linter will ignore them, so it will not throw error

	// Error
	// const data string
	// data = "123"
	// also you cannot declare constant with :=
	// data0 := 123 // golang will recognize this as variable not constant
	// const data0_ := 123 // this is will error

	// not error
	const data1 = "123"
	// or
	const data2 string = "123"

	// declare multiple constant

	const (
		data3      = 123
		data4      = "456"
		data5 bool = false
	)

}
