package main

import "fmt"

func main() {
	var tmp float64
	fmt.Print("Enter Fahrenheit temp: ")
	fmt.Scan(&tmp)
	fmt.Println((tmp-float64(32))*5/9, "degrees Celsius")
}
