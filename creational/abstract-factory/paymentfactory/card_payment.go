package paymentfactory

import (
	"design-patterns/creational/abstract-factory/paymentmethod"
	"errors"
)

type CardPaymentFactory struct {
}

func (c *CardPaymentFactory) Get(cardType int) (Payment, error) {
	switch cardType {
	case Credit:
		return new(paymentmethod.CreditCardPayment), nil
	case Debit:
		return new(paymentmethod.DebitCardPayment), nil
	default:
		return nil, errors.New("not supported")
	}
}
