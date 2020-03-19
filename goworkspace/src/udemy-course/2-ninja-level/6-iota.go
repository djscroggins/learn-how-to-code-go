package main

import "fmt"

const (
	thisYear = 2020
	a        = thisYear + iota
	b        = thisYear + iota
	c        = thisYear + iota
	d        = thisYear + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
