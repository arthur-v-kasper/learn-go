package main

import (
	"errors"
	"fmt"
)

//Its possivel create a new type
type Bitcoin int

//Wallet type returns a wallet struct
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in teste is %v \n", &w.balance)
	return w.balance
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsuficintFunds = errors.New("cannot withdraw, isufficient funds =/")

// Withdraw subtracts some Bitcoin from the wallet, returning an error if it cannot be performed.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsuficintFunds
	}

	w.balance -= amount
	return nil
}
