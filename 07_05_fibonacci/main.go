package main

import "fmt"

func main() {
	fmt.Println(fibonacci(6))
}

func fibonacci(x int) int {
	var ret int
	if x<=1 {
		ret = x
	} else {
		ret = fibonacci(x-1) + fibonacci(x-2)
	}
	return ret
}
