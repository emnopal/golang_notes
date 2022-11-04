package main

import (
	"belajar_go/basic/23_import/helper"
	_ "belajar_go/basic/23_import/other" // ini gak akan dipanggil di fungsi disini, tapi dia otomatis dipanggil karena ada init() di dalam package other buat hindarin error linter golang jadinya harus ditambahin _ (blank)
	"fmt"
)

func main() {
	// function init blm di eksekusi
	// krn dia bakalan di ekseskusi otomatis
	// ketika helper dijalankan
	greetings := helper.SayGreeting()
	fmt.Println(greetings)

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
