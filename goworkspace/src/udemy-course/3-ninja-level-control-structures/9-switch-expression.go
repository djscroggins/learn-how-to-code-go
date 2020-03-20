package main

import "fmt"

func main() {
	favSport := "baseball"
	msg := "My favorite sport is"

	switch favSport {
	case "baseball":
		fmt.Printf("%s %s\n", msg, favSport)
	case "football":
		fmt.Printf("%s %s\n", msg, favSport)
	case "hockey":
		fmt.Printf("%s %s\n", msg, favSport)
	case "soccer":
		fmt.Printf("%s %s\n", msg, favSport)
	case "basketball":
		fmt.Printf("%s %s\n", msg, favSport)

	}
}
