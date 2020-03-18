package main

import "fmt"

type myType2 int

var m2 myType2
var i int

func main() {
	fmt.Printf("m = %d\n", m2)
	fmt.Printf("m is of Type: %T\n", m2)
	m2 = 42
	fmt.Printf("m = %d\n", m2)
	i = int(m2)
	fmt.Printf("i = %d\n", i)
	fmt.Printf("m is of Type: %T\n", i)
}
