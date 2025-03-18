package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var errInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// package fmt uses Stringer interface and have a String method to output values in string with %s
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
