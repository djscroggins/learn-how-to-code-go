package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) speak(){
	fmt.Printf("Hello, my name is %s and I am %d years old\n", p.firstName, p.age)
}

func main() {
	p := Person{
		firstName: "David",
		lastName:  "Scroggins",
		age:       37,
	}

	p.speak()
}
