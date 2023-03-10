package main

import (
	"fmt"
	"testing"
)

type Student struct {
	Name        string
	Height      float64
	Age         int32
	IsGraduated bool
	Hobbies     []string
}

func initializeStudent() (data *Student) {
	data = &Student{
		Name:        "Lorem",
		Height:      172.12,
		Age:         24,
		IsGraduated: true,
		Hobbies:     []string{"sport", "music"},
	}
	return
}

// note: for str layout only works for fmt.Printf()

func TestNumericStrLayout(t *testing.T) {
	var data *Student = initializeStudent()

	// normal numeric as it should
	fmt.Printf("%d\n", data.Age)

	// numeric to bytes
	fmt.Printf("%b\n", data.Age) // return binary ==> 11000

	// numeric to hex
	fmt.Printf("%x\n", data.Age) // return binary ==> 18
	// or
	fmt.Printf("%X\n", data.Age) // for uppercase; return binary ==> 18

	// numeric to octal
	fmt.Printf("%o\n", data.Age) // return octal ==> 30

	// bytes to numeric
	fmt.Printf("%d\n", 0b101011) // return human number ==> 43

	// hex to numeric
	fmt.Printf("%d\n", 0xFFA121) // return human number ==> 16752929

	// octal to numeric
	fmt.Printf("%d\n", 0o13343) // 5859

	// float formating
	// %e (or %E) => scientific
	// %f (or %.<num:int>f or %F or %.<num:int>F) => float formating
	// %g (or %.<num:int>g or %G or %.<num:int>G) => decimal formating

	// scientific notation (must be float)
	fmt.Printf("%e\n", data.Height) // 1.721200e+02
	// or
	fmt.Printf("%E\n", data.Height) // 1.721200E+02

	// formating float number (must be float)
	fmt.Printf("%f\n", data.Height)    // 172.120000
	fmt.Printf("%.0f\n", data.Height)  // 172
	fmt.Printf("%.1f\n", data.Height)  // 172.1
	fmt.Printf("%.2f\n", data.Height)  // 172.12
	fmt.Printf("%.32f\n", data.Height) // 172.12000000000000454747350886464119

	// decimal numeric (must be float and better representation if 1 > num > -1)
	fmt.Printf("%g\n", data.Height)   // 172.12
	fmt.Printf("%.0g\n", data.Height) // 2e+02 <- rounding
	fmt.Printf("%.1g\n", data.Height) // 2e+02 <- rounding
	fmt.Printf("%.2g\n", data.Height) // 1.7e+02 <- rounding
	fmt.Printf("%.3g\n", data.Height) // 172 <- rounding
	fmt.Printf("%.4g\n", data.Height) // 172.1 <- rounding

	// decimal numeric (must be float and better representation if 1 > num > -1); better example
	fmt.Printf("%g\n", data.Height/123456)   // 0.0013941809227579057
	fmt.Printf("%.0g\n", data.Height/123456) // 0.001 <- rounding
	fmt.Printf("%.1g\n", data.Height/123456) // 0.001 <- rounding
	fmt.Printf("%.2g\n", data.Height/123456) // 0.0014 <- rounding
	fmt.Printf("%.3g\n", data.Height/123456) // 0.00139 <- rounding
	fmt.Printf("%.4g\n", data.Height/123456) // 0.001394 <- rounding

	// with scientific notation
	fmt.Printf("%E\n", data.Height/123456) // 1.394181E-03

}

func TestPointerStrLayout(t *testing.T) {
	var data *Student = initializeStudent()

	// print pointer address
	fmt.Printf("%p\n", &data) // 0x<hex_of_your_pointer_location>
}

func TestEscapeStr(t *testing.T) {
	// print with escape character
	fmt.Printf("%q\n", `" name \ height "`) // will print someting like => \" name \\ height \"
}

func TestFormatStr(t *testing.T) {

	var data *Student = initializeStudent()

	// print without escape character
	fmt.Printf("%s\n", `" name \ height "`) // will print someting like => " name \ height ";

	// other uses

	// print standard string
	fmt.Printf("%s\n", "foo")                                               // foo
	fmt.Printf("%s, %s, %s\n", data.Name, data.Hobbies[0], data.Hobbies[1]) // Lorem, sport, music

}

func TestUnicodeStrLayout(t *testing.T) {
	var unicodeChar int = 1322
	// unicode/numeric to string
	fmt.Printf("%c\n", unicodeChar) // return string ==> Ԫ
}

func TestBoolStrLayout(t *testing.T) {
	var data *Student = initializeStudent()

	fmt.Printf("%t\n", data.IsGraduated)
}

func TestTypeStrLayout(t *testing.T) {
	var data *Student = initializeStudent()

	// print to get type of variable
	fmt.Printf("%T\n", *data)
	fmt.Printf("%T\n", &data)
	fmt.Printf("%T\n", data)
	fmt.Printf("%T\n", data.Age)
	fmt.Printf("%T\n", data.Height)
	fmt.Printf("%T\n", data.Hobbies)
	fmt.Printf("%T\n", data.IsGraduated)
}

func TestStructStrLayout(t *testing.T) {
	var data *Student = initializeStudent()

	// print struct object
	fmt.Printf("%#v\n", data)
}
