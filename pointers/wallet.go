package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) Stringer() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//struct pointers are automatically dereferenced
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("Oh no!")
	}
	w.balance -= amount
	return nil
}

// do not require struct pointer as paramter, it is just for consistency.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
