package main

import (
	"fmt"
	"math"
)

func square(num float64) float64 {
	return math.Pow(num, 2)
}

func main() {
	localSquare := square
	result := localSquare(100)
	fmt.Println(result)
}
