package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("expected %s got %s", want, got)
		}
	})
	t.Run("Test withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("expected %s got %s", want, got)
		}
	})
	t.Run("Test withdraw insufficient founds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("expected %s got %s", want, got)
		}

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

}
