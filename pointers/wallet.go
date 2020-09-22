package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// 在 Go 中，当调用一个函数或方法时，参数会被复制
// 当调用这个方法时，w 来自我们调用这个方法的副本
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is ", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

/**
指针
	当你传值给函数或方法时，Go 会复制这些值。如果需要更改状态，就需要用到指针指向你想要更改的值
*/
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}
