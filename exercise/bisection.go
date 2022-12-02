package main

import "exercise/pkg"

func main() {
	result := pkg.Bisection{}
	result.CalculateBisection(2, 4, true)
}
