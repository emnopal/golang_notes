package pkg

func Pic(dx, dy int) [][]uint8 {

	slices := make([][]uint8, dx)
	for x := range slices {
		slices[x] = make([]uint8, dy)
		for y := range slices[x] {
			slices[x][y] = TwoVarParabolicFunction(float64(x), float64(y))
		}
	}

	return slices
}
