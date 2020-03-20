package main

import "fmt"

func main() {
	Container := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "David",
		lastName:  "Scroggins",
		age:       37,
	}

	fmt.Println(Container)
}
