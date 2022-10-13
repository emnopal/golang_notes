package main

import "fmt"

func main() {

	// note: all of their data inside array must be same data type
	// array size cannot increased after array have been initialize

	// initialize array with assign in one by one
	var test_array [4]string
	test_array[0] = "test1"
	test_array[1] = "test2"
	test_array[2] = "test3"
	test_array[3] = "test4"
	// test_array[4] = "test5" // this is not okay, since the size of array is 4, this is overflow

	fmt.Println(test_array)      // get all array data
	fmt.Println(test_array[3])   // get an array data with index
	fmt.Println(test_array[1:3]) // get some of array data with slicing

	// initialize array without assign in one by one
	var test_array2 [4]string = [4]string{"test1", "test2", "test3", "test4"}

	fmt.Println(test_array2)      // get all array data
	fmt.Println(test_array2[3])   // get an array data with index
	fmt.Println(test_array2[1:3]) // get some of array data with slicing

	// or
	test_array3 := [4]string{
		"test1",
		"test2",
		"test3", // this is okay, since the size of array is 4, so no overflow
	}

	// or
	// var test_array3 = 4[string]{"test1", "test2", "test3"}

	fmt.Println(test_array3) // get all array data

	// get an array data with index
	fmt.Println(test_array3[3]) // print nothing since the size of array is only 3

	fmt.Println(test_array3[1:3]) // get some of array data with slicing

	// or with ellipsis

	// ellipsis, this is also best practice
	test_array4 := [...]string{"test1", "test2", "test3", "test4", "test5"}
	fmt.Println(test_array4)      // get all array data
	fmt.Println(test_array4[3])   // get an array data with index
	fmt.Println(test_array4[1:3]) // get some of array data with slicing

	// array with specifying index
	test_array5 := [...]string{1: "test1", 3: "test2", 2: "test3", 5: "test4", 6: "test5", 0: "0", 4: "0"}
	fmt.Println(test_array5)      // get all array data
	fmt.Println(test_array5[3])   // get an array data with index
	fmt.Println(test_array5[1:3]) // get some of array data with slicing

	// get length of array
	fmt.Println(len(test_array5))

	test_6 := [6]string{}
	fmt.Println(len(test_6)) // this will return length of size array, not how many data inside array; result=6

	// change element of array
	test_array5[4] = "new_test"
	fmt.Println(test_array5)
	fmt.Println(test_array5[4])

	// array declaration with make
	array_with_make := make([]int, 3) // []int: type, 3: capacity, length => capacity == length
	fmt.Println(len(array_with_make))
	fmt.Println(cap(array_with_make))
	fmt.Println(len(array_with_make) == cap(array_with_make))

	// multidimensional array
	multidim_array := [2][3]int{{3, 2, 3}, {3, 4, 5}} // 2 rows, 3 cols
	fmt.Println(multidim_array)

}
