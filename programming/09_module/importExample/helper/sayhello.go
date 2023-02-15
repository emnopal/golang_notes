package helper

import "fmt"

type People struct {
	Name, Nationality string
}

// pkg initialization
// diakses ketika package ini dijalankan

var greet string

func init() {
	greet = "Welcome...\n"
}

func SayGreeting() string {
	return greet
}

func SayHelloInEnglish(people People) {
	fmt.Println("Hello", people.Name)
}

func (people *People) SayHello() {
	if people.Nationality == "Indonesia" {
		fmt.Println("Halo", people.Name)
	} else if people.Nationality == "Japan" {
		fmt.Println("こんにちは", people.Name)
	} else {
		fmt.Println("Hello", people.Name)
	}
}
