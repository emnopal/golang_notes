package main

import "fmt"

func main() {

	// switch case
	name := "test"
	switch name {
	case "test":
		fmt.Println("name not valid")
	case "foo":
	case "bar":
		fmt.Println("name foo and bar is not valid")
	default:
		fmt.Println("Hello", name)
	}

	// short statement
	switch length := len(name); length < 5 || length > 30 {
	case true:
		fmt.Println("too long or too short")
	case false:
		fmt.Println("okay")
	}

	// switch without condition
	length_name := len(name)
	switch { // no condition, equivalent with if else
	case length_name < 5:
		fmt.Println("Too short")
	case length_name > 30:
		fmt.Println("Too long")
	default:
		fmt.Println("Okay")
	}
}
