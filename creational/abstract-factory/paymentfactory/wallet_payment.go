package paymentfactory

import (
	"design-patterns/creational/abstract-factory/paymentmethod"
	"errors"
)

type WalletPaymentFactory struct {
}

func (w *WalletPaymentFactory) Get(walletType int) (Payment, error) {
	switch walletType {
	case Paytm:
		return new(paymentmethod.PaytmWalletPayment), nil
	case Phonepe:
		return new(paymentmethod.PhonepeWalletPayment), nil
	default:
		return nil, errors.New("not supported")
	}
}
