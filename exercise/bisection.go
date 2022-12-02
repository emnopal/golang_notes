package main

import "exercise/pkg"

func main() {
	result := pkg.Bisection{}
	result.CalculateBisection(1, 0.001, true)
}
