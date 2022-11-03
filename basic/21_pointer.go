package main

import "fmt"

/*
pass by value?
secara default di golang, semua variable itu pass by value not reference
artinya, jika kita kirim variable ke dalam function, method atau variable lain
itu yg dikirim adalah duplikasi dari value nya atau artinya kita gak bener2 akses si memori address nya
melainkan kopian dari memori address nya, jadi ketika ada perubahan maka memori address dari value yg
aslinya tidak berubah

kalau pass by reference itu kita secara directly akses ke memori address nya
jadi apa yg kita assign itu lah nilai aslinya, artinya tidak di kopi
sehingga jika ada perubahan maka memori address nya yg akan di replace

simple nya
pass by value: c++, c, java, go, rust
pass by reference: python, javascript, php

pass by value: value dari sebuah variable jika diassign / dikirim ke variable lain
				itu bukan value aslinya tapi kopian value nya

pass by reference: value dari sebuah variable jika diassign / dikirim ke variable lain
				itu value aslinya

kemudian apa itu pointer?
nah pointer itu adalah tools di bahasa pemrograman yg pass by value
buat ngereference/nge"point" (knp namanya pointer) variable ke letak memory address aslinya
jadi ngebuat si variable jadi pass by reference
*/

type Address struct {
	city     string
	province string
	region   string
	country  string
}

func main() {

	// pass by value and their problem
	address1 := Address{
		city:     "Jakarta Utara",
		province: "DKI Jakarta",
		region:   "Jabodetabek",
		country:  "Indonesia",
	}
	address2 := address1
	address2.city = "Jakarta Selatan"
	fmt.Println(address1) // ini gak berubah, city nya tetep jakarta utara
	fmt.Println(address2) // ini berubah jadi jakarta selatan
	// nah klo di bahasa yg pass by reference kyk php, js atau python
	// harusnya si address1 city nya berubah juga ke jakarta selatan
	// nah disini problem nya, jadi si address2 itu cuma ambil value dari address1
	// lalu di copy ke alamat memori yg lain, sehingga value di address1 itu tetap
	// beda kaya di bahasa pass by reference
	// klo itu datanya berubah di address1
	// klo di pass by value supaya data yg di assign ke variable lain berubah jg itu harus pake pointer (* (asterix))
	// klo di pass by reference supaya data yg di assign ke variable lain tetap itu harus pake copy (python) atau & (other)

	// how to fix pass by value problem and make it pass by reference
	// using pointer (pass by value to pass by reference (use *), pass by reference to pass by value (use &) -> only in golang)
}
