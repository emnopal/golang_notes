package helper

import "fmt"

type People struct {
	Name, Nationality string
}

func SayHelloInEnglish(people People) {
	fmt.Println("Hello", people.Name)
}

func (people *People) SayHello() {
	if people.Nationality == "Indonesia" {
		fmt.Println("Hi", people.Name)
	} else if people.Nationality == "Japan" {
		fmt.Println("こんにちは", people.Name)
	} else {
		fmt.Println("Hello", people.Name)
	}
}
