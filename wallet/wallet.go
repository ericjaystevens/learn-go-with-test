package wallet

import (
	"errors"
	"fmt"
)

type Wallet struct {
	ballance Bitcoin
}

type Bitcoin int

type Stringer interface {
	String() string
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount < w.GetBallance() {
		w.ballance -= amount
		return nil

	} else {
		return errors.New("Account does not have suffient funds")
	}
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) GetBallance() Bitcoin {
	return w.ballance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.ballance += amount
}
