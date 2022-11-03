package main

import (
	"errors"
	"fmt"
)

/*
error interface, mirip sama inherit Exception class di bahasa pemrograman lain

java/c++/python/js
class CustomException extends Exception{}

golang
type error interface {
	Error() string
}
atau simple nya tinggal pake aja package errors

error beda dengan panic
kalo error dia gak akan break code, jadi cukup simple handling dengan conditional
kalo panic dia bakalan break code, jadi harus ada handling
*/

func divided(numerator int, denumerator int) (float64, error) {
	if denumerator == 0 {
		return 0, errors.New("0 Denumerator") // harus 2 return value, karena return type nya ada 2, float dan error
		// nah yg float itu hasil, dan yg error itu error nya
		// tapi karena mau raise error, jadinya yg hasil nya di set 0, gabisa di set nil, karena bukan interface/function/map
	}
	return float64(numerator) / float64(denumerator), nil // sama kaya diatas, karena ada hasilnya dan error nya gaada jadi
	// error nya di set ke nil
}

func main() {
	result, error := divided(12, 0)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println(result)
	}
	result1, error1 := divided(12, 8)
	if error1 != nil {
		fmt.Println(error1.Error())
	} else {
		fmt.Println(result1)
	}
}
