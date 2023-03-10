package main

import "fmt"

func runAppsPanic(error bool) {
	if error {
		panic("Panic!")
	}
	fmt.Println("Not reachable line")
}

func runAppsPanicWithDefer(error bool) {
	defer func() {
		fmt.Println("Executed right before panic is executed!")
	}()
	if error {
		panic("Panic!")
	}
	fmt.Println("Not reachable line")
}

func main() {

	// panic
	// function yg bisa kita gunakan untuk menghentikan program
	// jadi kita buat error buatan
	// mirip kaya raise Exception klo di python atau throw Exception bahasa lain
	// runAppsPanic(true) // next line never be executed, solve this with recover()
	runAppsPanic(false) // next line will be executed

	// defer akan tetap di jalankan walaupun terjadi panic
	// runAppsPanicWithDefer(true) // next line never be executed, solve this with recover()
	runAppsPanicWithDefer(false) // next line will be executed

	// so next statement can be executed
	fmt.Println("Never executed!!!")

}
