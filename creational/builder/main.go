package main

import (
	"design-patterns/creational/builder/farecalculation"
	"fmt"
)

type RideBooking struct {
	vehicleType         string
	sourceLocation      float64
	destinationLocation float64
}

func main() {
	fmt.Printf("-------------------Builder Pattern--------------\n\n")
	rides := getRides()
	for _, ride := range rides {
		fmt.Printf("Calculating fare for the %s ride\n", ride.vehicleType)
		taxi := farecalculation.Taxi{}
		fare, err := taxi.Calculate(ride.vehicleType, ride.sourceLocation, ride.destinationLocation)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Fare without Tax:%d\nFare with Tax:%d\n", fare.FareWithoutTax, fare.FareWithTax)
		}
		fmt.Println("-----------------------------------------------")
	}

}

func getRides() []RideBooking {
	rides := []RideBooking{
		{vehicleType: "bike", sourceLocation: 10.42, destinationLocation: 25.61},
		{vehicleType: "car", sourceLocation: 15.75, destinationLocation: 34.43},
		{vehicleType: "bus", sourceLocation: 12.86, destinationLocation: 28.98},
	}

	return rides
}
