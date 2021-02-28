package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		//fmt.Printf("address of balance in teste is %v \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsuficintFunds)

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	//t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	//t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one =(") //t.Fatal stop execution
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
