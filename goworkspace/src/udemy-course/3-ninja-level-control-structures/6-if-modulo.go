package main

import "fmt"

func main() {
	for i := 0; i <= 10_000; i++ {
		mod := i % 4

		if mod == 0 {
			fmt.Printf("%d is divisble by 4\n", i)
		}
	}
}
