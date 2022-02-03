package paymentfactory

import "errors"

const (
	Card    = 1
	Wallet  = 2
	Debit   = 3
	Credit  = 4
	Paytm   = 5
	Phonepe = 6
)

type PaymentFactory interface {
	Get(int) (Payment, error)
}

type Payment interface {
	Pay()
}

func GetPaymentFactory(paymentType int) (PaymentFactory, error) {
	switch paymentType {
	case Card:
		return new(CardPaymentFactory), nil
	case Wallet:
		return new(WalletPaymentFactory), nil
	default:
		return nil, errors.New("not supported")
	}
}
