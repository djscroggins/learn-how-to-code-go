package main

import "fmt"

func main() {
	_slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%T", _slice)

	for _, v := range _slice {
		fmt.Println(v)
	}
}
