package main

import "fmt"

/*
di golang, function itu bisa berdiri sendiri alias istilahnya first citizen class
gak kaya di java, function harus ada di dalam class
function juga merupakan tipe data, ini mirip di python juga, krn function jg object di python, tp klo di golang engga ya
function juga bisa disimpan di variable, mirip di python juga
*/

// function
func simpleFunc() {
	fmt.Println("Hello")
}

// function with parameter
func bitComplicateFunc(name string) {
	fmt.Println("Hello:", name)
}

func moreComplicateFunc(name string, age int8) { // go doesn't support default parameter
	fmt.Println("Hello:", name, "Age:", age)
}

// function returning a value
func bitComplicateFuncButWithReturnValue(name string) string {
	return "Name: " + name
}

// function returning multiple value
func moreComplicateButWithReturnValue(name string, age int8) (string, int8) {
	return "Name: " + name, age
}

// if parameter's type is same
func ifParamsIsSame(name, nationality string) (string, string) {
	return name, nationality
}

// named return value
// for (name, nationality, address string) this is only work if all of return value type is equal
// (all of the 3 is string), if not you have to write the types e.g. (name string, age int8)
func namedReturnValue() (name, nationality, address string) {
	// no need to reassign if type is defined, keyword ':=' is assign
	name = "John"
	nationality = "Japan"
	address = "Tokyo, Japan"
	return // not need to declare return value since it's been declared in named return value
}

func namedReturnValueIfTypeIsDifference() (name string, age int8) {
	name = "test"
	age = 23
	return // not need to declare return value since it's been declared in named return value
}

// variadic function
// mirip args kwargs di python
// bisa terima lebih dari input yg di define
// jadi dia pake array buat nerima datanya
// tapi gaperlu declare array, cukup pisahin pake koma

func sumAll(numbers ...int) int { // gabisa 2 variadic cuma bisa satu dan cuma bisa di paling akhir/kanan/final/belakang
	// func test(a ...string, b string) // error
	// func test(b string, a ...string) // bisa
	total := 0
	isNumExists := false // klo mau deklarasi pake :=, gabisa di redeclared lagi
	for _, number := range numbers {
		total += number
		isNumExists = true // klo mau ganti value cukup =
	}

	if isNumExists {
		return total
	}

	return 0
}

// variadic function juga bisa dipake buat cek parameter atau cek data di array
func checkIsExists(params ...string) bool {
	for _, par := range params {
		if par == "test" {
			return false
		}
	}
	return true
}

// variadic function with other arguments
func mergeAllNumsWithOps(ops string, nums ...int) int { // variadic harus di akhir semua argument
	isNumExists := false
	switch ops {
	case "+":
		total := 0
		for _, number := range nums {
			total += number
			isNumExists = true
		}
		if isNumExists {
			return total
		}
		return 0
	case "-":
		total := 0
		for _, number := range nums {
			total -= number
			isNumExists = true
		}
		if isNumExists {
			return total
		}
		return 0
	case "*":
		total := 1
		for _, number := range nums {
			total *= number
			isNumExists = true
		}
		if isNumExists {
			return total
		}
		return 0
	case "/":
		total := 1
		for _, number := range nums {
			total /= number
			isNumExists = true
		}
		if isNumExists {
			return total
		}
		return 0
	default:
		return 0
	}
}

// function as value, simpan function di variable
func getConst(const_name string) float32 {
	switch const_name {
	case "pi":
		return 3.14159265359
	case "e":
		return 2.71828
	default:
		return 0
	}
}

// function as parameter
func countAreaOfCircle(nums float32, const_nums func(string) float32) float32 {
	return nums * const_nums("pi")
}

// function as parameter
// you can use alias to make it shorter
type ConstNums func(string) float32

func countConst(nums float32, const_nums ConstNums, constant string) float32 {
	return nums * const_nums(constant)
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursion

func factorialNonRecursion(nums uint) uint {
	var total uint = 1
	var i uint
	for i = 1; i <= nums; i++ {
		total *= i
	}
	return total
}

func factorialRecursion(nums uint) uint {
	if nums <= 1 {
		return 1
	}
	return nums * factorialRecursion(nums-1)
}

// note for recursion: make sure not overflow

func main() {
	simpleFunc()
	bitComplicateFunc("Test")
	moreComplicateFunc("Test", 23)

	resultReturnValue := bitComplicateFuncButWithReturnValue("Budi")
	fmt.Println(resultReturnValue)

	// for returning multiple value, you have to explicitly declared them
	// test := moreComplicateButWithReturnValue("str", 33) // will error, since it only declare one value
	name, age := moreComplicateButWithReturnValue("Budi", 28)
	fmt.Println(name)
	fmt.Println(age)

	// if you don't want to declare them
	_, age2 := moreComplicateButWithReturnValue("Joko", 33)
	fmt.Println(age2)

	_, nationality, _ := namedReturnValue()
	fmt.Println(nationality)

	// variadic function
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	// memasukan slice atau array ke dalam variadic function
	nums := []int{10, 20, 30, 40}
	total_new := sumAll(nums...) // tambahin ... dibelakang slice
	fmt.Println(total_new)

	// bisa juga kirim data kosongan ke variadic function
	fmt.Println(sumAll())

	// panggil checkIsExists
	fmt.Println("is Exists?: ", checkIsExists("dev", "test"))

	// initialize variadic function with other arguments
	fmt.Println(mergeAllNumsWithOps("/", 1, 2, 3, 4, 5))

	// initialize function as value
	math_const := getConst
	fmt.Println(math_const("pi"))
	// or
	math_const_e := getConst("e")
	fmt.Println(math_const_e)

	// function as parameter
	math_const_pi := getConst
	fmt.Println(countAreaOfCircle(12, math_const_pi))

	// or
	another_const := getConst
	fmt.Println(countConst(12, another_const, "e"))

	// anonymous function; lambda function in python
	cubic_nums := func(nums int) int {
		return nums * nums * nums
	}
	fmt.Println(cubic_nums(5))

	// assign anonymous function into function
	fmt.Println(countConst(12, func(string) float32 {
		return 3.14159265359
	}, "e")) // force const to pi
	// or

	// assign function as value
	fmt.Println(countConst(float32(cubic_nums(5)), func(string) float32 {
		return 3.14159265359
	}, "e"))
	// or
	cub_num := cubic_nums(10)
	fmt.Println(countConst(float32(cub_num), func(string) float32 {
		return 3.14159265359
	}, "e"))

	// this is error
	// fmt.Println(countConst(func() float32 { // cannot use func() type to type float32
	// 	return 12
	// }, func(string) float32 {
	// 	return 3.14159265359
	// }, "e"))

	// another example of anonymous function

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("test", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("test", func(name string) bool {
		return name == "root"
	})

	// recursive function
	fmt.Println(factorialNonRecursion(6))
	fmt.Println(factorialRecursion(6))
}
