package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regexp.MustCompile(string) *regexp.Regexp
	// similar with re.compile in python
	pattern := regexp.MustCompile("(test)")
	fmt.Println(pattern)

	// regexp.MatchString(string) bool
	// if match return true else false
	fmt.Println(pattern.MatchString("testo"))       // true
	fmt.Println(pattern.MatchString("test"))        // true
	fmt.Println(pattern.MatchString("testa"))       // true
	fmt.Println(pattern.MatchString("testb"))       // true
	fmt.Println(pattern.MatchString("test string")) // true
	fmt.Println(pattern.MatchString("tjest"))       // true

	// regexp.FindAllString(string, max_array) bool
	// if substring match return the string which match
	fmt.Println(pattern.FindAllString("testo test testa testb tjest toest", -1)) // [test test test test]

}
