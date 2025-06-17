package main

import "fmt"

// Account
// struct representing a bank account
type Account struct {
	Name    string
	Balance float64
}

// ShowBalance
// Method with value receiver
func (acc Account) ShowBalance() {
	fmt.Println("Account Holder:", acc.Name, "\nBalance: ", acc.Balance)
}

// AddFunds
// Method to add funds in account
func (acc *Account) AddFunds(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	acc.Balance += amount
	fmt.Println("deposited", amount, "to account", acc.Name)
}

// DeductFunds
// Method to deduct funds from account
func (acc *Account) DeductFunds(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return
	}
	if amount > acc.Balance {
		fmt.Println("Insufficient funds")
		return
	}
	acc.Balance -= amount
	fmt.Println("withdrawn", amount, "from account", acc.Name)
}

func main() {
	// Create a new account
	acc := Account{Name: "Jatin", Balance: 100.0}
	fmt.Println("Welcome to the Bank!")

	acc.ShowBalance()

	// Deposit
	acc.AddFunds(200.0)
	acc.ShowBalance()

	// Withdraw
	acc.DeductFunds(150.0)
	acc.ShowBalance()

	// Modify
	accCopy := acc
	accCopy.AddFunds(1000.0) // Won't affect the original
	fmt.Println("Account modified:")
	accCopy.ShowBalance()

	fmt.Println("Original account unchanged:")
	acc.ShowBalance()
}
