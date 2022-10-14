package main

import "fmt"

// function
func simpleFunc() {
	fmt.Println("Hello")
}

// function with parameter
func bitComplicateFunc(name string) {
	fmt.Println("Hello:", name)
}

func moreComplicateFunc(name string, age int8) { // go doesn't support default parameter
	fmt.Println("Hello:", name, "Age:", age)
}

// function returning a value
func bitComplicateFuncButWithReturnValue(name string) string {
	return "Name: " + name
}

// function returning multiple value
func moreComplicateButWithReturnValue(name string, age int8) (string, int8) {
	return "Name: " + name, age
}

// if parameter's type is same
func ifParamsIsSame(name, nationality string) (string, string) {
	return name, nationality
}

// named return value
// for (name, nationality, address string) this is only work if all of return value type is equal
// (all of the 3 is string), if not you have to write the types e.g. (name string, age int8)
func namedReturnValue() (name, nationality, address string) {
	// no need to reassign if type is defined, keyword ':=' is assign
	name = "John"
	nationality = "Japan"
	address = "Tokyo, Japan"
	return // not need to declare return value since it's been declared in named return value
}

func namedReturnValueIfTypeIsDifference() (name string, age int8) {
	name = "test"
	age = 23
	return // not need to declare return value since it's been declared in named return value
}

func main() {
	simpleFunc()
	bitComplicateFunc("Test")
	moreComplicateFunc("Test", 23)

	resultReturnValue := bitComplicateFuncButWithReturnValue("Budi")
	fmt.Println(resultReturnValue)

	// for returning multiple value, you have to explicitly declared them
	// test := moreComplicateButWithReturnValue("str", 33) // will error, since it only declare one value
	name, age := moreComplicateButWithReturnValue("Budi", 28)
	fmt.Println(name)
	fmt.Println(age)

	// if you don't want to declare them
	_, age2 := moreComplicateButWithReturnValue("Joko", 33)
	fmt.Println(age2)

	_, nationality, _ := namedReturnValue()
	fmt.Println(nationality)

}
