package main

import (
	"fmt"
	"reflect"
)

/*
Reflection adalah teknik untuk inspeksi variabel,
mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
Cakupan informasi yang bisa didapatkan lewat reflection sangat luas,
seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

2 fungsi yg sering dipake: reflect.ValueOf() dan reflect.TypeOf().

Fungsi (reflect.ValueOf() -> reflect.Value): informasi yang berhubungan dengan nilai pada variabel yang dicari
Fungsi (reflect.TypeOf() -> reflect.Type): informasi yang berhubungan dengan tipe data variabel yang dicari
*/

/*
Reflect bisa digunakan untuk mengambil informasi semua property
variabel objek cetakan struct, dengan catatan property-property tersebut bermodifier public.
*/

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr { // <- reflect.Ptr is struct
		reflectValue = reflectValue.Elem() // get all element in struct
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name) // get name
		fmt.Println("tipe data :", reflectType.Field(i).Type) // get type
		//fmt.Println("nilai     :", reflectValue.Field(i).Interface()) // get value; ok, but better with this:
		switch value := reflectValue.Field(i).Interface(); result := value.(type) {
		case string:
			fmt.Println("nilai     :", result)
		case bool:
			fmt.Println("nilai     :", result)
		case int:
			fmt.Println("nilai     :", result)
		case float64:
			fmt.Println("nilai     :", result)
		default:
			fmt.Println("nilai     :", result)
		}
		fmt.Println("")
	}
}

/*
Informasi mengenai method juga bisa diakses lewat reflect,
syaratnya masih sama seperti pada pengaksesan proprerty, yaitu harus bermodifier public.
*/

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	number := 4
	reflectValue := reflect.ValueOf(number)

	fmt.Println(reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println(reflectValue.Int())
	}

	// return empty interface aka any
	fmt.Println(reflectValue.Interface())

	// cast any to int (or other type)
	// using switch case to avoid error
	switch result := reflectValue.Interface(); value := result.(type) {
	case string:
		fmt.Println(value)
	case bool:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	case float64:
		fmt.Println(value)
	default:
		fmt.Println(value)
	}

	// accessing struct member info using reflect
	fmt.Println()
	s1 := &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()

	// accessing method info using reflect
	s2 := &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s2.Name)

	reflectValueS2 := reflect.ValueOf(s2)
	methodReflectValueS2 := reflectValueS2.MethodByName("SetName")
	methodReflectValueS2.Call([]reflect.Value{
		reflect.ValueOf("wick john"),
	})

	fmt.Println("nama :", s2.Name)
}
