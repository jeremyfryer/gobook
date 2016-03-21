package main

import (
	"fmt"
	m "github.com/jeremyfryer/gobook/11_packages/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println("The average is:", avg)

	fmt.Println("The max is:", m.Max(xs))
	fmt.Println("The min is:", m.Min(xs))

}
