package main

import "fmt"

func main() {

	/*
		type declaration is making alias with existing data type
	*/

	type yearStr string
	type yearInt int16

	var todayYear yearStr = "2022"
	var todayYearInt32 int32 = 2022
	var todayYearInt yearInt = yearInt(todayYearInt32)

	fmt.Println(todayYear)
	fmt.Println(todayYearInt)

}
