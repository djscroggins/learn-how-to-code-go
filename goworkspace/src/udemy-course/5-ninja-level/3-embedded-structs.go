package main

import "fmt"

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	vehicle   Vehicle
	fourWheel bool
}

type Sedan struct {
	vehicle Vehicle
	luxury  bool
}

func main() {

	truck := Truck{
		vehicle: Vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	sedan := Sedan{
		vehicle: Vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(truck)
	fmt.Println(truck.vehicle.doors)
	fmt.Println(sedan)
	fmt.Println(sedan.luxury)

}
