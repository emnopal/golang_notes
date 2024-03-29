package main

import "fmt"

/*
struct itu kumpulan nol atau lebih dari tipe data
istilahnya buat bikin banyak custom data type
yg gabisa dilakukan di type alias
struct juga bisa buat type balikan

tapi selain itu, struct biasanya representasi
data spt data pada array, json dsb, karena array itu cuma bisa satu type doang
nah klo struct dia bisa banyak type

map juga sebenernya bisa juga, tapi key:value nya itu harus sama, klo satu key dia string
semua, maka harus string semua, tipe key sm value bisa beda, tapi tiap type itu, klo
kita deklarasiin lagi itu harus sama

contoh

{
	"a": 2,
	3: "c"
} // ini error, krn tipe nya gak sama

data di struct bisa disimpan di field

kesimpulan: struct adalah kumpulan dari field

atau anggap aja deh struct ini class, biar lebih gampang :p
*/

// keyword struct itu sama kaya type alias, cuma tambahin keyword struct aja setelah nama_field nya
type Customer struct {
	Name, Address string
	Age           uint8
	Member        bool
}

// NOTE: struct cuma bisa add field aja, gabisa tambahin function di dalam struct
// klo mau buat function, maka harus buat method

// struct method
// mirip method di class
func (customer Customer) isPremium() {
	if customer.Member {
		fmt.Println("Yes", customer.Name, "is Premium")
	} else {
		fmt.Println("No", customer.Name, "is not Premium")
	}
}

// cara lain
// cuma yg ini gak rekomen
func isAgeOK(customer Customer) {
	if customer.Age > 20 {
		fmt.Println("Yes", customer.Name, "is OK")
	} else {
		fmt.Println("No", customer.Name, "is not OK")
	}
}

// embedded struct
type person struct {
	name string
	age  int
}

// Embed struct juga bisa dibilang inheritance di golang, karena akan manggil
// semua inherited method/property yg diambil dari struct yg lain
type student struct {
	grade  int
	person // bisa ambil dari struct lain
}

func main() {

	// cara menggunakan struct, atau klo di oop buat object

	// best practice pake ini daripada harus pake array atau map klo mau simpen
	// kecuali klo kompleks baru harus pake gabungan antara array map dan struct
	var member_1 Customer

	member_1.Address = "Tokyo"
	member_1.Name = "John"
	member_1.Age = 32
	member_1.Member = false

	fmt.Println(member_1)
	fmt.Println(member_1.Age)
	member_1.isPremium()

	// cara lain deklarasi struct: struct literal
	member_2 := Customer{
		Name:    "Jacob",
		Address: "Osaka",
		Age:     45,
		Member:  true,
	}

	fmt.Println(member_2)
	member_2.isPremium()

	// buat cara lain/cara ke2, harus masukin struct ke function nya
	// tapi ini gak di rekomendasikan
	// karena gak keliatan spt object
	isAgeOK(member_1)
	isAgeOK(member_2)

	// embedded struct
	student_1 := student{}
	student_1.name = "wick"
	student_1.age = 21
	student_1.grade = 2

	fmt.Println(student_1) // {2 {wick 21}}
	// since struct student contains {grade person{}}
	// so the result would be {grade {name, age}}

}
