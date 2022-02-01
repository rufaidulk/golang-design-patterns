package main

import (
	"design-patterns/creational/factory-method/faresettingsfactory"
	"fmt"
)

func main() {
	fareSetting, err := faresettingsfactory.GetFareSetting(faresettingsfactory.Car)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fareSetting.Find())
	}
}
