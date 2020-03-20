package main

import "fmt"

func main() {
	countYear := 1982
	currentYear := 2020
	for countYear <= currentYear {
		fmt.Println(countYear)
		countYear++
	}
}
