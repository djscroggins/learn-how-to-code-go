package main

import "fmt"

func returnInt() int64 {
	return 64
}

func returnIntAndString() (string, int64) {
	return "sixty four", 128
}

func main(){
	i := returnInt()
	fmt.Printf("i = %d\n", i)
	s, i := returnIntAndString()
	fmt.Printf("s = %s; i = %d\n", s, i)
}
