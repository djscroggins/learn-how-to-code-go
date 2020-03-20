package main

import "fmt"

func main() {
	for decimalEncoding := 65; decimalEncoding <= 90; decimalEncoding++ {
		fmt.Println(decimalEncoding)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", decimalEncoding)
		}
	}
}
