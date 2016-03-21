package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

/*
[]float64 -> float64
Finds the minimum of a slice of float64s
 */
func Min(xs []float64) float64 {
	min := xs[0]
	for _, value := range xs {
		if value < min {
			min = value
		}
	}
	return min
}


/*
[]float64 -> float64
Finds the maximum of a slice of float64s
 */
func Max(xs []float64) float64 {
	max := xs[0]
	for _, value := range xs {
		if value > max {
			max = value
		}
	}
	return max
}
