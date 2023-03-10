package main

import (
	"belajar_go/basic/23_import/example"
	. "belajar_go/basic/23_import/helper" // ada titik disini, supaya jadi selevel, jadi kita gaperlu ketik helper.People{} lagi, cukup People{} aja
	_ "belajar_go/basic/23_import/other"  // ini gak akan dipanggil di fungsi disini, tapi dia otomatis dipanggil karena ada init() di dalam package other buat hindarin error linter golang jadinya harus ditambahin _ (blank)
	"fmt"
)

func main() {

	// defer juga bisa berfungsi sebagai destructor
	defer example.End()

	// function init blm di eksekusi
	// krn dia bakalan di ekseskusi otomatis
	// ketika helper dijalankan
	greetings := SayGreeting()
	fmt.Println(greetings)

	// people0 := helper.People{"Adam", "UK"} // without using key is bad practice

	people1 := People{
		Name:        "John",
		Nationality: "US",
	}
	people2 := People{
		Name:        "Budi",
		Nationality: "Indonesia",
	}
	people3 := People{
		Name:        "Miyama",
		Nationality: "Japan",
	}
	SayHelloInEnglish(people1)
	people2.SayHello()
	people3.SayHello()
}
