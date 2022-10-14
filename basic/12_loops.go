package main

import "fmt"

func main() {
	// the only loop in golang is for loop

	// ordinary for loop golang
	// mirip while loop di c++
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println()

	// another example of loops
	// mirip for loop di c++
	for initCounter := 0; initCounter <= 10; initCounter++ {
		fmt.Println(initCounter)
	}

	// iterating array or slice
	fmt.Println()

	arrExample := []int{10, 20, 30, 40, 50, 60}

	for i := 0; i < len(arrExample); i++ {
		fmt.Println(arrExample[i])
	}

	// for range
	// mirip loop di python
	fmt.Println()

	for index, nums := range arrExample {
		fmt.Println("Index:", index, "Nums:", nums)
	}

	for _, nums := range arrExample { // _ for disabling index, since golang is strict if variable is not used
		fmt.Println("Nums:", nums)
	}

	fmt.Println()
	// iterating map
	person := map[string]string{
		"name":        "John Doe",
		"address":     "USA",
		"nationality": "USA",
	}

	for key, value := range person {
		fmt.Println("Key:", key, "Value:", value)
	}

	// infinite loop, just for example
	// sums := 0
	// for {
	// 	fmt.Println(sums)
	// 	sums++ // infinite loop
	// }
	fmt.Println()

	// breaking loop
	sums := 0
	for {
		fmt.Println(sums)
		if sums == 10 {
			break
		}
		sums++ // infinite loop
	}

	fmt.Println()

	// continue loop
	for _, str := range "iterating to string" {
		if string(str) == " " {
			continue // ignoring whitespace
		} else {
			fmt.Println(string(str)) // cast back from byte to string
		}
	}

}
