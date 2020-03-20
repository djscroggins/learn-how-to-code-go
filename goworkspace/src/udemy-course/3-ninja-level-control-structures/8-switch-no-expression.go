package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This line won't print")
	case true:
		fmt.Println("This line will print")
	}
}
