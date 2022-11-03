package main

import "fmt"

func main() {
	// map is key-value (dictionary in python)
	// key must be unique
	// size of map is unlimited but the key must be unique
	// if we use same key (key not unique), the value will be replaced with new value

	// map example
	person := map[string]string{
		"name":        "John Doe",
		"address":     "USA",
		"nationality": "USA",
	}

	fmt.Println(person)         // by default, golang will return map key alphabetically so the result will be addres, name and nationality
	fmt.Println(person["name"]) // in golang, to access map key, must used "", if use '' will be error
	fmt.Println(person["address"])
	fmt.Println(person["nationality"])

	// replace map value with key
	person["address"] = "Tokyo, Japan"

	// add key value to map
	person["job"] = "Programmer"
	person["title"] = "Bachelor"

	fmt.Println(person)

	// delete map key:value
	delete(person, "title")

	fmt.Println(person)

	// create empty map
	newMap := make(map[string]string)

	fmt.Println(newMap)

}
