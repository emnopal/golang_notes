package main

import "fmt"

/*
secara default, object yg blm di inisialisasi maka secara default dia nilainya nil, null atau None
tapi beda klo di golang, kalau di golang saat kita buat variable dgn tipe data tertentu,
maka secara default variable tersebut ada nilainya.
contoh:
int/float -> default nya 0
bool -> default nya false
string -> default nya ""
struct -> default nya field kosong

di golang juga ada tipe data nil, nil sama kaya null atau None
tapi nil hanya bisa digunakan di interface, function, map, slice, pointer dan channel
*/

// contoh nil
// di function
func newMap(name string) map[string]string {
	if name == "" {
		return nil // dia return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {

	// nil map
	var nilMap map[string]string = nil
	fmt.Println(nilMap)
	if nilMap == nil {
		fmt.Println("Nil map")
	}

	// nil function
	newMapNoNil := newMap("John")
	fmt.Println(newMapNoNil) // map[name: John]
	if newMapNil := newMap(""); newMapNil == nil {
		fmt.Println(newMapNil) // map[]
		fmt.Println("Nil function")
	}

}
