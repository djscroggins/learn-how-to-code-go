package main

import "fmt"

func main() {
	customers := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for customer, data := range customers{
		fmt.Printf("\nCutomer = %s\n-----------------------\n", customer)
		for _, datum := range data {
			fmt.Println(datum)
		}
		fmt.Println("-----------------------")

	}
}
