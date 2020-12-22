package wallet

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	t.Run("deposit funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		assertBalance(t, wallet, Bitcoin(10))

		wallet.Deposit(100)
		assertBalance(t, wallet, Bitcoin(110))
	})

	t.Run("empty withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(1)}
		err := wallet.Withdraw(Bitcoin(1))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("full withdrawal", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("insufficent funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(1000))

		assertBalance(t, wallet, Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("invalid withdrawal amount", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(-10))

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, ErrNegativeWithdraw)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	noError := got == nil
	wrongError := got != want

	switch {
	case noError:
		t.Fatal("expected an error and didn't get one")
	case wrongError:
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an err but did not want one")
	}
}
