package main

import "fmt"

func main() {
	_array := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T", _array)

	for _, v := range _array {
		fmt.Println(v)
	}
}
