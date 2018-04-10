package main

import "fmt"

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

func vectorCross(u, v []float64) (res []float64) {
	res = make([]float64, len(u))
	// < AyBz - AzBy , AzBx - AxBz , AxBy - AyBy >
	res[0] = u[1]*v[2] - u[2]*v[1]
	res[1] = u[2]*v[0] - u[0]*u[2]
	res[2] = u[0]*v[1] - u[1]*v[1]
	return
}

func vectorDot(u, v []float64) float64 {
	fmt.Println(u)
	fmt.Println(v)
	output := 0.0
	for i, _ := range u {
		output += u[i] * v[i]
	}
	return output
}
