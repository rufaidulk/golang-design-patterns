package farecalculation

import "design-patterns/creational/builder/faresettings"

type CarFareCalculator struct {
	fare        Fare
	fareSetting faresettings.FareSettings
}

func (c *CarFareCalculator) init(fareSetting faresettings.FareSettings) TaxiFareCalculator {
	c.fareSetting = fareSetting

	return c
}

func (c *CarFareCalculator) calculateFare(sourceLocation float64, destinationLocation float64) TaxiFareCalculator {
	distance := c.findDistanceTravelled(sourceLocation, destinationLocation)
	c.findBaseFare(distance)
	c.addServiceCharge()

	return c
}

func (c *CarFareCalculator) getFare() Fare {
	return c.fare
}

func (c *CarFareCalculator) findDistanceTravelled(sourceLocation float64, destinationLocation float64) int {
	return int(destinationLocation - sourceLocation)
}

func (c *CarFareCalculator) findBaseFare(distance int) {
	c.fare.FareWithoutTax = distance * c.fareSetting.BaseFare
	c.fare.FareWithTax = c.fare.FareWithoutTax + c.fareSetting.Tax
}

func (c *CarFareCalculator) addServiceCharge() {
	c.fare.FareWithoutTax += c.fareSetting.ServiceCharge
	c.fare.FareWithTax += c.fareSetting.ServiceCharge
}
