package main

import "fmt"

/*
Interface itu tipe data abstract, jadi tidak diimplementasikan secara langsung
interface itu isinya definisi method atau property
interface dipake sbg kontrak

nah klo mau gunain kontrak/interface nya harus deklarasiin struct yg sama data nya dgn interface
*/

// harus di implementasikan
type HasName interface {
	GetName() string // tinggal di implementasikan
}

// cara implementasi nya
type Person struct {
	name string
}

// buat implementasi GetName(), tinggal panggil aja fungsi nya ke struct
func (person Person) GetName() string {
	return person.name
}

// panggil fungsi yg implementasi GetName()
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// walaupun struct nya beda, tapi masih bisa diimplementasikan
// dan gaakan bertabrakan sm struct Person
type Animal struct {
	name string
}

func (animal Animal) GetName() string {
	return animal.name
}

func main() {
	member := Person{
		name: "Dani",
	}
	sayHello(member)
	pet := Person{
		name: "Cat",
	}
	sayHello(pet)
}
