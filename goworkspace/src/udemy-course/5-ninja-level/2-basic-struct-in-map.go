package main

import "fmt"

type Person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func summarizePerson(person Person) {
	fmt.Printf("\nName: %s %s\n", person.firstName, person.lastName)
	fmt.Println("Favorite Ice Cream Flavors:")
	for i, flavor := range person.favoriteIceCreamFlavors {
		fmt.Printf("  %d. %s\n", i+1, flavor)
	}
}

func main() {
	p1 := Person{
		firstName:               "David",
		lastName:                "Scroggins",
		favoriteIceCreamFlavors: []string{"Salted Caramel", "Mint Chocolate Chip"},
	}

	p2 := Person{
		firstName:               "James",
		lastName:                "Connor",
		favoriteIceCreamFlavors: []string{"Fudge", "Cookie Dough", "Neapolitan"},
	}

	people := map[string]Person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range people {
		summarizePerson(v)
	}

}
