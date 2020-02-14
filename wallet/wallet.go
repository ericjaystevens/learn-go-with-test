package wallet

import(
)

type Wallet struct{
	ballance Bitcoin
}

type Bitcoin int

func (w *Wallet) GetBallance() Bitcoin{
	return w.ballance
}

func (w *Wallet) Deposit(amount Bitcoin){
	w.ballance += amount
}