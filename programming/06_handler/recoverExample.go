package main

import "fmt"

// jadi buat catch error nya harus di defer
// cara 1
// buat function baru
func recoverWrapper() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
}

func runAppsRecover(error bool) {
	defer func() {
		fmt.Println("Panic detected!")
	}()
	defer recoverWrapper()
	if error {
		panic("Error!")
	}
	// this will not executed
	fmt.Println("App running")
}

// =======================================================

// atau bisa seperti ini
// tambahin defer di existing function yg mau di defer
func endAppWithRecover() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
	fmt.Println("End endAppWithRecover")
}

func runAnotherAppsRecover(error bool) {
	defer endAppWithRecover() // pastikan recover() diekseskusi di akhir
	if error {
		panic("Error!")
	}
	// this will not executed
	fmt.Println("App running")
}

// =======================================================

// cara lain
// buat recover di dalam scope function
func anotherRecoverCase() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	num := 0
	total := 12 / num
	// this will not executed
	fmt.Println(total)

}

// =======================================================

// cara lain
// buat recover di dalam scope function
func anotherRecoverCase2() {
	defer func() {
		if message := recover(); message != nil {
			fmt.Println(message)
		}
	}()

	num := 0
	total := 12 / num
	// this will not executed
	fmt.Println(total)

}

func main() {

	// recover
	// function yg bisa kita gunakan untuk catch panic
	// sehingga program tidak terhenti ketika ada panic
	// mirip kaya try except di python atau try catch di bahasa lain
	runAppsRecover(true) // not breaking a code
	// contoh lain
	runAnotherAppsRecover(true) // not breaking a code
	// contoh lain
	anotherRecoverCase()
	// contoh lain
	anotherRecoverCase2()

	// so next statement can be executed
	fmt.Println("Executed! Finished!!!")

}
