package main

import "belajar_go/basic/23_import/helper"

func main() {
	people1 := helper.People{
		Name:        "John",
		Nationality: "US",
	}
	people2 := helper.People{
		Name:        "Budi",
		Nationality: "Indonesia",
	}
	people3 := helper.People{
		Name:        "Miyama",
		Nationality: "Japan",
	}
	helper.SayHelloInEnglish(people1)
	people2.SayHello()
	people3.SayHello()
}
