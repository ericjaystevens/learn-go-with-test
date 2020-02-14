package wallet

import(
	"testing"
)

func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.GetBallance()
	want := Bitcoin(10)

	if got != want{
		t.Errorf("got: %d wanted: %d", got, want)
	}
}