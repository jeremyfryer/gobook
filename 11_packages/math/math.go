package math

//Average finds the average of a series of numbers
func Average(xs []float64) float64 {
	if len(xs) > 0 {
		total := float64(0)
		for _, x := range xs {
			total += x
		}
		return total / float64(len(xs))
	}
	return 0

}

//Min finds the minimum of a slice of float64s
//[]float64 -> float64
func Min(xs []float64) float64 {
	min := xs[0]
	for _, value := range xs {
		if value < min {
			min = value
		}
	}
	return min
}

//Max finds the maximum of a slice of float64s
//[]float64 -> float64
func Max(xs []float64) float64 {
	max := xs[0]
	for _, value := range xs {
		if value > max {
			max = value
		}
	}
	return max
}
