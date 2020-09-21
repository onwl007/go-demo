package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	// 当我们在代码中更改 balance 的值时，我们处理的是来自测试的副本，因此，balance 在测试中没有被改变
	fmt.Println("address of balance in test is ", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
