package main

import "fmt"

func main() {
	/*
		slice adalah potongan dari array, dia merupakan tipe data reference
		slice mirip dengan array, tapi slice punya size yg fleksible
		slice dapat mengakses data di array dengan metode slicing
		data slice memiliki 3 tipe: pointer, length, capacity
		1. pointer adalah penunjuk data pertama di array
		2. length adalah panjang dari slice
		3. capacity adalah kapasitas dari slice dimana length !> dari capacity

		slicing array
		array[low:high] return slice from array[low] to array[high]
		array[low:] return slice from array[low] to end of array(n)
		array[:high] return slice from beginning(0) of array to array[high]
		array[:] return slice from beginning(0) until end of array(n)
	*/

	// example

	test_array := [...]string{"test1", "test2", "test3", "test4", "test5"}
	test_slice := test_array[1:3]                       // pointer: 1, length: 2, capacity: from pointer until end (5-1 = 4)
	fmt.Println("test_slice", test_slice)               // ["test2", "test3"]
	fmt.Println("test_slice capacity", cap(test_slice)) // capacity of slice: 4
	fmt.Println("test_slice length", len(test_slice))   // length of slice: 2

	// append slice
	/*
		Ketika jumlah elemen dan lebar kapasitas adalah sama (len(fruits) == cap(fruits)),
		maka elemen baru hasil append() merupakan referensi baru.

		Ketika jumlah elemen lebih kecil dibanding kapasitas (len(fruits) < cap(fruits)),
		elemen baru tersebut ditempatkan ke dalam cakupan kapasitas, menjadikan semua
		elemen slice lain yang referensi-nya sama akan berubah nilainya.
	*/

	// note, disini test_array is reference (ini bakalan keubah), test_slice<1,2,3,dst> itu slice baru
	fmt.Println()
	test_slice1 := append(test_slice, "foo_test")
	fmt.Println("test_slice1 (new_slice)", test_slice1) // slice baru: [test2 test3 foo_test]
	fmt.Println("test_slice (old_slice)", test_slice)   // slice baru: [test2 test3]
	fmt.Println(
		"test_array (old_array_reference_slice)",
		test_array,
	) // reference: [test1 test2 test3 foo_test test5], this will replace after test3 (from test_array[1:3])

	fmt.Println()
	test_slice2 := append(test_slice1, "bar_test")
	fmt.Println("test_slice2 (new_slice)", test_slice2) // slice baru: [test2 test3 foo_test bar_test]
	fmt.Println("test_slice1 (old_slice)", test_slice1) // slice baru: [test2 test3 foo_test]
	fmt.Println(
		"test_array (old_array_reference_slice)",
		test_array,
	) // reference: [test1 test2 test3 foo_test bar_test], this will replace after test3 (from test_array[1:3])

	fmt.Println()
	test_slice3 := append(test_slice2, "foobar_test")
	fmt.Println("test_slice3 (new_slice)", test_slice3) // slice baru: [test2 test3 foo_test bar_test foobar_test]
	fmt.Println("test_slice2 (old_slice)", test_slice2) // slice baru: [test2 test3 foo_test bar_test]
	fmt.Println(
		"test_array (old_array_reference_slice)",
		test_array,
	) // reference: [test1 test2 test3 foo_test bar_test], this will replace after test3 (from test_array[1:3])
	// the value after test_array[1:3] is ignored (foobar_test)

	// create new slice from scratch without array
	new_slice := make([]string, 4, 6) // []string: type, 4: length, 6: capacity
	fmt.Println(len(new_slice))       // 4
	fmt.Println(cap(new_slice))       // 6

	// copy slice
	test_array2 := []string{"test1", "test2", "test3", "test4", "test5"}
	make_array2_slice := make([]string, 3, 3) // []string: type, 3: length, 3: capacity
	copy(make_array2_slice, test_array2)
	fmt.Println(make_array2_slice)

	// caution, jgn sampe salah buat array dan slice
	this_is_array := [...]string{"a"}
	this_is_slice := []string{"a"}

	fmt.Println(this_is_array)
	fmt.Println(this_is_slice)

	// akses slice dengan 3 element

	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[1:2]   // 1: pointer, 2: length
	var bFruits = fruits[1:2:3] // 1: pointer, 2: length, 3: capacity

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 3

}
