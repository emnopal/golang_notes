package main

import "fmt"

func finishingRunning() {
	fmt.Println("This will be executed even if the function is panic")
}

// example for running defer in normal function
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

func runAppsPanic(execute_panic bool) {
	defer finishingRunning() // defer harus di awal supaya jika ada error, gak ngeblok execute function nya
	if execute_panic {
		panic("panic executed")
	}
	fmt.Println("this is not reachable line")
}

func main() {

	// defer
	// defer adalah function yg bisa dijadwalkan untuk di ekseskusi ketika
	// ada sebuah function yg selesai di eksekusi walaupun terjadi error
	// dia tetap akan di ekseskusi
	// mirip finally
	// atau mirip juga destructor

	runApps()
	// runAppsError() // still executed but then error thrown
	// runAppsPanic(true) // still executed but then error thrown

}
