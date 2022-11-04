package main

import (
	"fmt"
	s "strings"
)

// learn more: golang.org/pkg/strings

func main() {

	// contoh beberapa
	// Trim
	strTrim := s.Trim("!!!!!Hello, World!!!!!", "!")
	fmt.Println(strTrim)

	// toLower
	strToLower := s.ToLower("HEllo, WoRlD")
	fmt.Println(strToLower)

	// toUpper
	strToUpper := s.ToUpper("hello, World")
	fmt.Println(strToUpper)

	// split
	strSplit := s.Split("Hello World", " ")
	fmt.Println(strSplit)

	// contains
	strContains := s.Contains("Hello:World", ":")
	fmt.Println(strContains)

	// replace
	strRpl := s.ReplaceAll("Hello World", "o", "a") // apa bedanya dgn strings.Replace
	fmt.Println(strRpl)
}
