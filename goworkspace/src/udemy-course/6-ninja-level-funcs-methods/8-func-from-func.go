package main

import (
	"fmt"
	"math"
)

func square(num float64) float64 {
	return math.Pow(num, 2)
}

func getSquareFunc() func(float64) float64 {
	return square
}

func main() {
	localSquare := getSquareFunc()
	result := localSquare(100)
	fmt.Println(result)
}
