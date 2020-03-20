package main

import "fmt"

func main() {
	for i := 0; i <= 10_000; i++ {
		msg := "is divisible by"
		mod3 := i % 3
		mod4 := i % 4

		if mod3 == 0 {
			fmt.Printf("%d %s 3\n", i, msg)
		} else if mod4 == 0 {
			fmt.Printf("%d %s 4\n", i, msg)
		} else {
			fmt.Printf("%d is not divisble by 3 or 4\n", i)
		}
	}
}
