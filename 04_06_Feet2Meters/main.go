package main

import "fmt"

func main() {
	var feet float64
	fmt.Print("Enter number of feet: ")
	fmt.Scan(&feet)
	fmt.Println(feet*0.3048, "meters")
}
