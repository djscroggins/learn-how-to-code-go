package main

import "fmt"

func variadicSum(integers ...int) int {
	sum := 0
	for _, i := range integers {
		sum += i
	}
	return sum
}

func typedSum(integers []int) int {
	sum := 0
	for _, i := range integers {
		sum += i
	}
	return sum
}

func main() {
	integers := []int{12, 532, 189, 738, 12, 1478390}
	sum1 := variadicSum(integers...)
	fmt.Println(sum1)
	sum2 := typedSum(integers)
	fmt.Println(sum2)

}
