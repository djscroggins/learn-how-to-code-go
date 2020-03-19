package main

import "fmt"

var a int = 25
var b int = 33

func main() {
	c := (a == b)
	d := (a <= b)
	e := (a >= b)
	f := (a != b)
	g := (a < b)
	h := (a > b)
	fmt.Println(c, d, e, f, g, h)
}
