package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("x: %v\n", x)
	y := append(x[:2], x[6:]...)
	fmt.Printf("y: %v\n", y)
}