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

// interface kosong
/*
golang bukan oop
jadi biasanya klo di bahasa oop itu ada superclass
contoh:
java ada java.lang.Object
python ada object/any/None/Any
javascript ada Object/object/any
nah buat adain si super class ini di golang
tinggal deklarasikan aja interface kosong
jadi dia akan otomatis membuat tipe data jadi implementasinya

contoh penggunaan interface kosong

fmt.Println(a ...interface{}) // bebas masukin tipe data apapun ke dalam Println
panic(v interface{}) // bebas masukin tipe data apapun ke dalam panic
recover() interface[] // return apapun ke fungsi recover()

*/

// interface kosong atau any

func sayInMultipleType() interface{} { // jadi any
	// return 2
	// return "test"
	return true // dia bisa return apa aja, tapi gabisa return multiple value
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

	testInterfaceKosong := sayInMultipleType()
	fmt.Println(testInterfaceKosong)

	// atau
	// var testInterfaceKosongLagi int = sayInMultipleType() // ini error, krn return type nya itu interface bkn int
	// solusinya
	var testInterfaceKosongLagi interface{} = sayInMultipleType()
	fmt.Println(testInterfaceKosongLagi)
	// atau
}
