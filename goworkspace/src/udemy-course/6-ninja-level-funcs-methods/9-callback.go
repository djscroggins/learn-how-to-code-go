package main

import (
	"fmt"
	"math"
)

func square(num int) int {
	numFloat := float64(num)
	squared := math.Pow(numFloat, 2)
	return int(squared)
}

func incrementThenConvert(stop int, convert func(int) int) int {
	value := 0
	for i := 0; i < stop; i++ {
		value++
	}
	return convert(value)
}

func main() {
	result := incrementThenConvert(2_012, square)
	fmt.Println(result)
}
