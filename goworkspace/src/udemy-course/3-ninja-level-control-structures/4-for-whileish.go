package main

import "fmt"

func main() {
	countYear := 1982
	currentYear := 2020
	for {
		fmt.Println(countYear)
		countYear++
		if countYear <= currentYear {
			continue
		} else {
			break
		}
	}
}
