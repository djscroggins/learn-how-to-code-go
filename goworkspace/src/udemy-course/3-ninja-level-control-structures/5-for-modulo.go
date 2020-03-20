package main

import "fmt"

func main() {
	for i := 11; i < 100; i++ {
		fmt.Printf("%d %% 4 = %d\n", i, i%4)
	}
}
