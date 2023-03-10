package main

import "fmt"

// defer panic recover mirip try catch finally di bahasa pemrograman lain

func finishingRunning() {
	fmt.Println("Finishing")
	fmt.Println("Finished")

}

func runApps() {
	defer finishingRunning() // will executing after runApps() executed
	fmt.Println("Start Executing")
	loop := 0
	for {
		if loop == 100 {
			break
		}
		loop++
	}
	fmt.Println("Executed")
}

func runAppsError() {
	defer finishingRunning() // defer harus di awal supaya jika ada error, gak ngeblok execute function nya
	nums := 0
	total := 12 / nums
	fmt.Println(total)
}

// panic
func endApp() {
	fmt.Println("End App")
}

func runAppsPanic(error bool) {
	defer endApp()
	if error {
		panic("Error!")
	}
	fmt.Println("this code is not reachable")
}

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
	defer endApp()
	defer recoverWrapper()
	if error {
		panic("Error!")
	}
	// this will not executed
	fmt.Println("App running")
}

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

func main() {

	// Defer     Panic        Recover; Equivalent:
	// Finally   Raise/Throw  Try-Catch/Try-Except

	// defer
	// defer adalah function yg bisa dijadwalkan untuk di ekseskusi ketika
	// ada sebuah function yg selesai di eksekusi walaupun terjadi error
	// dia tetap akan di ekseskusi
	// mirip finally
	// atau mirip juga destructor
	runApps()
	// runAppsError() // still executed but then error thrown

	// panic
	// function yg bisa kita gunakan untuk menghentikan program
	// jadi kita buat error buatan
	// mirip kaya raise Exception klo di python atau throw Exception bahasa lain
	// defer akan tetap di jalankan walaupun terjadi panic
	// runAppsPanic(true) // throw error
	runAppsPanic(false) // success

	// recover
	// function yg bisa kita gunakan untuk catch panic
	// sehingga program tidak terhenti ketika ada panic
	// mirip kaya try except di python atau try catch di bahasa lain
	runAppsRecover(true) // not breaking a code
	// contoh lain
	runAnotherAppsRecover(true) // not breaking a code
	// contoh lain
	anotherRecoverCase()

	// so next statement can be executed
	fmt.Println("Executed! Finished!!!")

}
