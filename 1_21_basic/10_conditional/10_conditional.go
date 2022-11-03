package main

import "fmt"

func main() {
	// example statement
	name := "test"
	age := false
	// address := "Jakata"

	if name == "test" {
		fmt.Println("Name wrong!")
	} else if name == "foo" {
		fmt.Println("Name wrong!")
	} else {
		fmt.Println("Hello")
	}

	if age {
		fmt.Println("Okay")
	}

	// if short statement
	if length_name := len(name); length_name < 5 || length_name > 30 {
		fmt.Println("Nama tidak valid: length is", length_name)
	}

	// if address { // cannot use non boolean expression (not like other language),
	// // so you must convert it to boolean or use comparation operator
	// 	fmt.Println("Okay")
	// }

}
