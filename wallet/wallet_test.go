package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBallance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()

		if w.GetBallance() != want {
			t.Errorf("got: %s wanted: %s", w.GetBallance(), want)
		}
	}
	t.Run("deposit and ballance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBallance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw and ballance", func(t *testing.T) {
		wallet := Wallet{ballance: 10}
		wallet.Withdraw(Bitcoin(5))
		assertBallance(t, wallet, Bitcoin(5))
	})
	t.Run("overdraw", func(t *testing.T) {
		wallet := Wallet{ballance: 5}
		e := wallet.Withdraw(Bitcoin(10))
		if e == nil {
			t.Errorf("Overdraw Error should have been returned, current ballance: %s", wallet.GetBallance())
		}
	})
}
