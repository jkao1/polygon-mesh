package main

func vectorSubtract(u, v []float64) (res []float64) {
	if len(u) != len(v) {
		return
	}
	res = make([]float64, len(u))
	for i := range u {
		res[i] = u[i] - v[i]
	}
	return
}

func vectorDot(u, v []float64) float64 {
	output := 0.0
	for i, _ := range u {
		output += u[i] * v[i]
	}
	return output
}
