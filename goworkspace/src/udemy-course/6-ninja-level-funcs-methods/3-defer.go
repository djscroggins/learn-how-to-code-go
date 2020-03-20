package main

import "fmt"

func deferDemo(stop int, bump int) (accumulator int) {
	defer func() {accumulator += bump}()
	for i := 0; i < stop ; i++  {
		accumulator++
	}
	return accumulator
}


func main() {
	result := deferDemo(100, 2000)
	fmt.Println(result)
	result = deferDemo(20, 47)
	fmt.Println(result)
}
