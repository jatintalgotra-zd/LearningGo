package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// PaymentMethod
// interface with method Pay(float64) string
type PaymentMethod interface {
	Pay(float64) string
}

// CreditCard
// struct defining CreditCard details and otp
type CreditCard struct {
	Card string
	otp  int
}

// Pay
// method to simulate payment
func (cred *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("%v : Paid %v using Card ending in %v\n", cred, amount, cred.Card)
}

// String
// method String to implement Stringer interface
func (cred *CreditCard) String() string {
	return "[CreditCard]"
}

// PayPal
// struct defining PayPal account details and otp
type PayPal struct {
	email string
}

// Pay
// method to simulate payment
func (pp *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("%v : Paid %v using PayPal account%v\n", pp, amount, pp.email)
}

// String
// method String to implement Stringer interface
func (pp *PayPal) String() string {
	return "[PayPal]"
}

// UPI
// struct defining UPI account details and otp
type UPI struct {
	uipID string
	otp   int
}

// Pay
// method to simulate payment
func (upi *UPI) Pay(amount float64) string {
	return fmt.Sprintf("%v : Paid %v using UPI: %v\n", upi, amount, upi.uipID)
}

// String
// method String to implement Stringer interface
func (upi *UPI) String() string {
	return "[UPI]"
}

// GenerateOTP
// function to simulate generating otp in supported methods using type switching
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
	// Payment Methods
	cc := CreditCard{Card: "4321"}
	pp := PayPal{"example@zop.dev"}
	upi := UPI{uipID: "example@okaxis"}

	// Slice of PaymentMethods to loop over
	slice := []PaymentMethod{&cc, &pp, &upi}

	// Looping over different Methods of Payment
	for _, method := range slice {
		// Type Switching
		switch p := method.(type) {
		// Type - CreditCard
		case *CreditCard:
			err := GenerateOTP(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v : Otp sent to registered number.\n", p)
			}
			fmt.Println(p.Pay(500))

		// Type - PayPal
		case *PayPal:
			err := GenerateOTP(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v : Otp sent to registered number.\n", p)
			}
			fmt.Println(p.Pay(500))

		// Type - UPI
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
