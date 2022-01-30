package farecalculation

import "design-patterns/creational/builder/faresettings"

type BikeFareCalculator struct {
	fare        Fare
	fareSetting faresettings.FareSettings
}

func (b *BikeFareCalculator) init(fareSetting faresettings.FareSettings) TaxiFareCalculator {
	b.fareSetting = fareSetting

	return b
}

func (b *BikeFareCalculator) calculateFare(sourceLocation float64, destinationLocation float64) TaxiFareCalculator {
	distance := b.findDistanceTravelled(sourceLocation, destinationLocation)
	b.findBaseFare(distance)
	b.addServiceCharge()

	return b
}

func (b *BikeFareCalculator) getFare() Fare {
	return b.fare
}

func (b *BikeFareCalculator) findDistanceTravelled(sourceLocation float64, destinationLocation float64) int {
	return int(destinationLocation - sourceLocation)
}

func (b *BikeFareCalculator) findBaseFare(distance int) {
	b.fare.FareWithoutTax = distance * b.fareSetting.BaseFare
	b.fare.FareWithTax = b.fare.FareWithoutTax + b.fareSetting.Tax
}

func (b *BikeFareCalculator) addServiceCharge() {
	b.fare.FareWithoutTax += b.fareSetting.ServiceCharge
	b.fare.FareWithTax += b.fareSetting.ServiceCharge
}
