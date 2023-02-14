package main

import (
	"fmt"
	"strconv"
)

func main() {

	// konversi string ke bool
	strToBool, err := strconv.ParseBool("t") // cuma bisa: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	if err == nil {
		fmt.Println(strToBool)
	} else {
		fmt.Println(err)
	}

	// konversi string ke integer
	// argument:
	// s: string,
	// base: 2 for "0b": binary, 8 for "0" or "0o": octal, 16 for "0x": hex, and 10: ordinary otherwise.
	// base: untuk parameter string nya depan nya harus "0<number>" contoh 0b11010 <- binary, 086743 <- octal, 0x983a4 <- hex, sisanya kaya angka biasa
	// bitSize: 0, 8, 16, 32, and 64
	// bitSize cuma ukuran di memori aja
	strToInt, err := strconv.ParseInt("-32", 10, 32)
	if err == nil {
		fmt.Println(strToInt)
	} else {
		fmt.Println(err)
	}
}
