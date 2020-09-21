package pointers

import "fmt"

type Wallet struct {
	balance int
}

// 在 Go 中，当调用一个函数或方法时，参数会被复制
// 当调用这个方法时，w 来自我们调用这个方法的副本
func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is ", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
