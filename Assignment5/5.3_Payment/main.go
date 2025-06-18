package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type PaymentMethod interface {
	Pay(float64) string
}

type CreditCard struct {
	Card string
	otp  int
}

func (cred *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("%v : Paid %v using Card ending in %v\n", cred, amount, cred.Card)
}
func (cred *CreditCard) String() string {
	return "[CreditCard]"
}

type PayPal struct {
	email string
}

func (pp *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("%v : Paid %v using PayPal account%v\n", pp, amount, pp.email)

}
func (pp *PayPal) String() string {
	return "[PayPal]"
}

type UPI struct {
	uipID string
	otp   int
}

func (upi *UPI) Pay(amount float64) string {
	return fmt.Sprintf("%v : Paid %v using UPI: %v\n", upi, amount, upi.uipID)
}
func (upi *UPI) String() string {
	return "[UPI]"
}

func GenerateOTP(p PaymentMethod) error {
	var err error
	switch pay := (p).(type) {
	case *CreditCard:

		pay.otp = rand.Intn(9000) + 1000

	case *UPI:
		pay.otp = rand.Intn(9000) + 1000

	default:
		err = errors.New("otp not supported")
	}
	return err
}
func main() {
	cc := CreditCard{Card: "4321"}
	pp := PayPal{"example@zop.dev"}
	upi := UPI{uipID: "example@okaxis"}

	slice := []PaymentMethod{&cc, &pp, &upi}

	for _, method := range slice {
		switch p := method.(type) {
		case *CreditCard:
			err := GenerateOTP(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v : Otp sent to registered number.\n", p)
			}
			fmt.Println(p.Pay(500))

		case *PayPal:
			err := GenerateOTP(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v : Otp sent to registered number.\n", p)
			}
			fmt.Println(p.Pay(500))

		case *UPI:
			err := GenerateOTP(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v : Otp sent to registered device\n", p)
			}
			fmt.Println(p.Pay(500))
		}
	}
}
