package main

import "fmt"

func main() {
	customers := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	newSlice := []string{"My catchphrase", "Pizza", "Long walks"}
	customers["added_record"] = newSlice

	for customer, data := range customers {
		fmt.Println("---------------------------------")
		fmt.Printf("Customer = %s\nData = %v\n", customer, data)

	}

	delete(customers, "bond_james")

	fmt.Printf("\nAFTER DELETE\n")
	for customer, data := range customers {
		fmt.Println("---------------------------------")
		fmt.Printf("Customer = %s\nData = %v\n", customer, data)

	}
}
