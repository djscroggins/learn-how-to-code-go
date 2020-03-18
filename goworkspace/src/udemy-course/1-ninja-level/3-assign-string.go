package main

import "fmt"

var _x int = 42
var _y string = "James Bond"
var _z bool = true

func main() {
	s := fmt.Sprintf("Default values: x = %d, y = %s, z = %t\n", _x, _y, _z)
	fmt.Println(s)
}
