package main

import "fmt"

func main() {
	_slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%T\n", _slice)

	a := _slice[:5]
	fmt.Println(a)

	b := _slice[5:]
	fmt.Println(b)

	c := _slice[2:7]
	fmt.Println(c)

	d := _slice[1:6]
	fmt.Println(d)
}
