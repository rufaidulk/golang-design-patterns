package farecalculation

import (
	"design-patterns/creational/builder/faresettings"
	"errors"
)

type Fare struct {
	FareWithoutTax int
	FareWithTax    int
}

type TaxiFareCalculator interface {
	init(faresettings.FareSettings) TaxiFareCalculator
	calculateFare(float64, float64) TaxiFareCalculator
	getFare() Fare
}

type Taxi struct {
	fareCalculator TaxiFareCalculator
}

func (t *Taxi) Calculate(vehicleType string, sourceLocation float64, destinationLocation float64) (fare Fare, err error) {
	switch vehicleType {
	case "car":
		t.fareCalculator = &CarFareCalculator{}
	case "bike":
		t.fareCalculator = &BikeFareCalculator{}
	default:
		err = errors.New("vehicle type is not supported")
	}

	if t.fareCalculator == nil {
		return fare, err
	}

	fare = t.fareCalculator.init(t.getFareSettings(vehicleType)).calculateFare(sourceLocation, destinationLocation).getFare()

	return fare, nil
}

func (t *Taxi) getFareSettings(vehicleType string) faresettings.FareSettings {
	//faresetting may be stored in DB
	return faresettings.FareSettings{
		BaseFare:      3,
		FarePerKm:     2,
		ServiceCharge: 2,
		Tax:           4,
	}
}
