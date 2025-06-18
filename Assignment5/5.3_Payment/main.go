package main

import (
	"fmt"
)

// PaymentMethod
// interface with method Pay(float64) string
type PaymentMethod interface {
	Pay(float64) string
}

// SupportsOTP
// interface with GenerateOTP method
type SupportsOTP interface {
	GenerateOTP()
}

// CreditCard
// struct defining CreditCard details and otp
type CreditCard struct {
	Card string
}

// GenerateOTP
// method to simulate generating OTP - to implement SupportsOTP interface
func (cred *CreditCard) GenerateOTP() {
	fmt.Printf("%v : Otp sent to registered number.\n", cred)
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

func (upi *UPI) GenerateOTP() {
	fmt.Printf("%v : OTP sent to registered device.\n", upi)
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

func main() {
	// Payment Methods
	cc := CreditCard{Card: "4321"}
	pp := PayPal{"example@zop.dev"}
	upi := UPI{uipID: "example@okaxis"}

	// Slice of PaymentMethods to loop over
	slice := []PaymentMethod{&cc, &pp, &upi}

	// Looping over different Methods of Payment
	for _, method := range slice {
		if otpSupported, ok := method.(SupportsOTP); ok {
			otpSupported.GenerateOTP()
		}
		method.Pay(500)
	}
}
