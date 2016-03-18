package main

import "fmt"

func main() {
	fmt.Println(greatest(10,30,50,40,20))
}

func greatest(x ...int) int {
	max := x[0]
	for _, value := range x {
		if value > max {
			max = value
		}
	}
	return max
}