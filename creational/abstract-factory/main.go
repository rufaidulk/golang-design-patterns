package main

import (
	"design-patterns/creational/abstract-factory/paymentfactory"
	"log"
)

const (
	Card    = 1
	Wallet  = 2
	Debit   = 3
	Credit  = 4
	Paytm   = 5
	Phonepe = 6
)

func main() {
	cardPaymentFactory, err := paymentfactory.GetPaymentFactory(Card)
	if err != nil {
		log.Fatal(err)
	}
	cardPayment, err := cardPaymentFactory.Get(Credit)
	if err != nil {
		log.Fatal(err)
	}
	cardPayment.Pay()

	walletPaymentFactory, err := paymentfactory.GetPaymentFactory(Wallet)
	if err != nil {
		log.Fatal(err)
	}
	walletPayment, err := walletPaymentFactory.Get(Paytm)
	if err != nil {
		log.Fatal(err)
	}
	walletPayment.Pay()
}
