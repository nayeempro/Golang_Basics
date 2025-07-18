package main

import "fmt"

type Account struct {
	balance int
}

// Pointer receiver allows modifying the actual value

func (a *Account) Deposit(amount int) {
	a.balance += amount
	fmt.Println("Deposit is ", amount)
}

// Show balance
func (a *Account) showBalance() {
	fmt.Println("Balance  is ", a.balance)
}

// withdraw
func (a *Account) Withdraw(amount int) {
	a.balance -= amount
	fmt.Println("withdraw amount is ", amount)
}

func main() {

	ac := Account{balance: 100}
	ac.showBalance()
	ac.Deposit(100)
	ac.showBalance()
	ac.Withdraw(50)
	ac.showBalance()
}
