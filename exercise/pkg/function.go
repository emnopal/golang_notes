package pkg

import "math"

func TwoVarParabolicFunction(x, y float64) uint8 {
	return uint8(math.Pow(x, 2) + math.Pow(y, 2))
}
