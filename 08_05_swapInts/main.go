package main

import "fmt"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
