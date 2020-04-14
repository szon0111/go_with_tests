package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("address of balance in Test is %v \n", &wallet.balance)

	want := Bitcoin(20)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
