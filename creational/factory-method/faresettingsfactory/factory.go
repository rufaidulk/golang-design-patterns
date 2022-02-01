package faresettingsfactory

import (
	"design-patterns/creational/factory-method/faresettings"
	"errors"
)

type VehicleType int8

const (
	Car VehicleType = 1 + iota
	Bike
	Bus
)

func GetFareSetting(vType VehicleType) (faresettings.FareSettings, error) {
	switch vType {
	case 1:
		return new(faresettings.CarFareSettings), nil
	case 2:
		return new(faresettings.BikeFareSettings), nil
	case 3:
		return new(faresettings.BusFareSettings), nil
	default:
		return nil, errors.New("vehicle type not supported")
	}
}
