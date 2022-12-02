package pkg

import (
	"fmt"
	"math"
)

type BisectionInterface interface {
	Function(x float64) float64
	DerivativeFunction(x float64) float64
}

type Bisection struct {
	BisectionInterface
}

func (bisection Bisection) Function(x float64) float64 {
	return 2*math.Pow(x, 4) + 4*math.Pow(x, 3) - 3*math.Pow(x, 2) - x + math.Cos(x) + 11
}

func (bisection Bisection) DerivativeFunction(x float64) float64 {
	return 8*math.Pow(x, 3) + 12*math.Pow(x, 2) - 6*x - 1 + math.Sin(x)
}

func (bisection Bisection) CalculateBisection(init float64, eps float64, debug bool) float64 {
	x := &init
	for math.Abs(init-*x) <= eps {
		funcx := bisection.Function(*x)
		dfuncx := bisection.DerivativeFunction(*x)
		init = *x - (funcx / dfuncx)
		if debug {
			fmt.Println(*x, init, math.Abs(init-*x))
		}
	}
	return init
}
