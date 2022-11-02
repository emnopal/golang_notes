package main

import "fmt"

/*
Closure adalah kemampuan function berinteraksi dgn data2 disekitar
dgn scope yg sama

Harap gunakan dengan bijak, karena fitur ini akan (baik sengaja maupun tidak sengaja)
mengubah data yg ada yg seharusnya kita tidak ubah

Scope di golang:
xxxx{
	// ini scope
}

// ini luar scope
*/

func main() {
	// fitur closure
	counter := 0 // dia bisa manggil ini
	increment := func() {
		fmt.Println("Increment ke-", counter)
		counter++ // manggil ini diluar scope increment
	}
	increment()          // increment ke-0
	increment()          // increment ke-1
	fmt.Println(counter) // manggil ini diluar scope increment; result = 2
	// hati hati karena bisa ubah data yg seharusnya gak diubah
	// karena dia bisa manggil sesuatu yg diluar scope nya

	// atau
	name := "test"
	fmt.Println(name)
	change_name := func() {
		name = "changed" // namanya keubah
	}
	change_name()
	fmt.Println(name)

	// jadi solusinya, usahain klo diluar scope (krn klo di dalam scope bakalan error)
	// di redeclared, atau pake variable lain biar gak ada variable shadow
}
