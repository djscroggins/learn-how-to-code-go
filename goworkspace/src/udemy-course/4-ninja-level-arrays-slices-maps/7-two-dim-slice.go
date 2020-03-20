package main

import "fmt"

func main() {
	customerOne := []string{"James", "Bond", "Shaken, not stirred"}
	customerTwo := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	customers := [][]string{customerOne, customerTwo}

	for _, customer := range customers {
		fmt.Println(customer)
		for _, datum := range customer {
			fmt.Println(datum)
		}
	}

}
