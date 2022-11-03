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
	city, province, region, country string
}

// pointer di function atau method
// by default, data yg dikirim ke function itu pass by value
// jadi hanya akan di proses di function saja
// nah gimana skrg klo pengen di ubah sampai ke luar function?
// pake pass by reference
// nah caranya gimana?
// caranya parameter nya diubah jadi pointer, tinggal tambahin operator * di parameter nya
// contoh tanpa pointer
func changeCountryWithoutPointer(address Address, country interface{}) {
	if country == nil {
		address.country = "Indonesia"
	} else {
		address.country = country.(string)
	}
}

// dengan pointer
func changeCountryWithPointer(address *Address, country interface{}) {
	if country == nil {
		address.country = "Indonesia"
	} else {
		address.country = country.(string)
	}
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
	// klo itu datanya berubah juga di address1

	// how to fix pass by value problem and make it pass by reference
	// using pointer!
	address3 := &address1 // tambahin amperstand, amperstand itu keyword reference
	address3.city = "Jakarta Barat"
	fmt.Println(address1) // city jadi jakarta barat
	fmt.Println(address2) // ini gak berubah karena di copy
	fmt.Println(address3) // city jadi jakarta barat, tapi pas di println ini kasitau klo address3 itu pointer

	fmt.Println()
	// nah klo mau eksplisit kasitau data type nya jadi begini
	var address4 *Address = &address1 // nah klo mau explisit tulis tipe datanya, harus ditambahin pointer di depan type nya
	address4.city = "Jakarta Timur"
	fmt.Println(address1) // city jadi jakarta timur, karena di reference sm address 4
	fmt.Println(address2) // ini gak berubah karena di copy
	fmt.Println(address3) // city jadi jakarta timur, karena sm-sm reference ke address1
	fmt.Println(address4) // city jadi jakarta timur

	fmt.Println()
	// tadi kan perubahan based nya cuma property dari struct nya
	// nah gimana klo ternyata struct nya di inisialisasi dari awal?
	address5 := Address{
		city:     "Jakarta Utara",
		province: "DKI Jakarta",
		region:   "Jabodetabek",
		country:  "Indonesia",
	}
	address6 := &address5

	// inisialisasi ulang
	address6 = &Address{
		city:     "Bekasi",
		province: "Jawa Barat",
		region:   "Jabodetabek",
		country:  "Indonesia",
	} // nah klo inisialisasi ulang harus ditambahin operator & juga karena address6 itu reference ke address5

	fmt.Println(address5) // ini gak berubah walaupun kita inisialisasi klo dia ini variable reference
	fmt.Println(address6)

	fmt.Println()
	// nah misalkan klo ada kasus si varibale ini harus berubah juga walaupun di inisialisasi ulang di variable reference baru?
	// caranya pake operator * (asteriks)
	address7 := Address{
		city:     "Jakarta Utara",
		province: "DKI Jakarta",
		region:   "Jabodetabek",
		country:  "Indonesia",
	}
	address8 := &address7

	*address8 = Address{
		city:     "Bogor",
		province: "Jawa Barat",
		region:   "Jabodetabek",
		country:  "Indonesia",
	} // kalo pake tanda *, gausah pake tanda & lagi
	// si address8 berubah jadi alamat memori yg utama karena pake tanda pointer * (asteriks)
	// sehingga si compiler nanti akan tunjuk si address8 ini sebagai reference dari address7
	// karena si address8 ini jadi alamat utama, sehingga nanti address7 yg harus reference kesini
	// jadi misalkan kalau kita buat variable baru
	address9 := &address7

	fmt.Println(address7) // berubah ikutin address8
	fmt.Println(address8) // tapi tetep di sini akan ada tanda & karena address8 itu variable hasil reference
	// walaupun reference nya berubah

	fmt.Println(address9) // tetep berubah jadi address8 walaupun reference nya ke address7,
	// krn istilahnya si address7 "dipaksa" reference ke address8 oleh operator *

	fmt.Println()
	// inisiaslisasi pointer kosong
	address10 := new(Address) // buat variable reference tanpa harus reference ke variable yg lain
	address11 := address10

	address11.city = "Tokyo"
	fmt.Println(address10)
	fmt.Println(address11) // otomatis reference ke address10

	fmt.Println()
	// atau
	var address12 *Address = new(Address)
	fmt.Println(address12) // jadi string kosong

	fmt.Println()
	// pointer in function
	// function tanpa pointer
	address13 := Address{
		city:     "Ginzha",
		province: "Tokyo",
		region:   "Kanto",
		country:  "",
	}
	changeCountryWithoutPointer(address13, "Japan")
	fmt.Println(address13) // gaakan keubah country nya ke Japan

	// gimana klo kita mau ubah?
	// solusinya pakein pointer
	changeCountryWithPointer(&address13, "Japan") // harus pake & karena address13 itu jadi pointer
	fmt.Println(address13)                        // berubah country nya jadi Japan

	// pointer juga berfungsi buat saving memory, karena jika banyak parameter yg di pass by value
	// maka akan semakin besar juga penggunaan memori nya
	// oleh karena itu, sebaiknya selalu gunakan pointer ketika ingin passing data ke function supaya
	// tidak terlalu banyak makan memori
	// apalagi jika struct nya besar sekali
}
