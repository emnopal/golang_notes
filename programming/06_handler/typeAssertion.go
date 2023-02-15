package main

import "fmt"

/*
type assertions: casting data type interface{} atau any
biasanya sering dipake klo tipe datanya any atau interface kosong
simpelnya ganti type interface{} atau any ke type yg kita mau
atau cating any atau interface{} ke string, int, bool dsb
*/

func isDepend(isTrue bool) interface{} {
	if isTrue {
		return "OK"
	}
	return false
}

func main() {
	// casting interface{} ke string
	resultTrue := isDepend(true) // "OK"
	resultAsStr := resultTrue.(string)
	fmt.Println(resultAsStr)

	// masih satu func, casting interface{} ke bool
	resultFalse := isDepend(false) // false
	resultAsBool := resultFalse.(bool)
	// resultAsBool := resultFalse.(string) // caution, this will cause panic() since this result is not string, but boolean
	fmt.Println(resultAsBool)

	// cara lebih aman supaya terhindar dari panic
	// jika kita malas nge-recover
	// ini lebih disarankan dan good practice

	switch result := isDepend(false); value := result.(type) {
	case string:
		fmt.Println(value, "with type is string")
	case bool:
		fmt.Println(value, "with type is bool")
	default: // ini klo kita gamau handle, misal int, float dsb
		fmt.Println(value, "is unknown type")
	}
}
