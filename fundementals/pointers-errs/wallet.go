package wallet

import (
	"errors"
)

// Bitcoin is what this wallet stores
type Bitcoin int

// Wallet struct returns data assiated with Wallet
type Wallet struct {
	balance Bitcoin
}

// Balance method returns wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit method adds value to the wallet
func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
}

// ErrInsufficientFunds returns when suffienct funds are not available for withdraw
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// ErrNegativeWithdraw returns when attempting to withdraw an ammount < 0
var ErrNegativeWithdraw = errors.New("cannot withdraw negative values")

// Withdraw lets users take money out of their accounts
func (w *Wallet) Withdraw(val Bitcoin) error {

	withdrawValueErr := val <= 0
	insufficentBalance := w.balance < val

	switch {
	case withdrawValueErr:
		return ErrNegativeWithdraw
	case insufficentBalance:
		return ErrInsufficientFunds
	default:
		w.balance -= val
		return nil
	}
}
