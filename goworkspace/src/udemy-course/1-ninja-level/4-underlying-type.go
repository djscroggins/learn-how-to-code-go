package main

import "fmt"

type myType int

var m myType

func main() {
	fmt.Printf("m = %d\n", m)
	fmt.Printf("m is of Type: %T\n", m)
	m = 42
	fmt.Printf("m = %d\n", m)
}
