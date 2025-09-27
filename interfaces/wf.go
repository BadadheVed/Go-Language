package main

type WellsFargo struct {
	balance int
}

func NewWellFargo() *WellsFargo {
	return &WellsFargo{
		balance: 0,
	}
}

func (w *WellsFargo) GetBalance() int {
	return w.balance
}

func (w *WellsFargo) Deposit(amount int) {
	w.balance += amount
}

func (w *WellsFargo) WithDraw(amount int) {
	if w.balance < amount {
		panic("Withdraw Failed")
	}
	w.balance -= amount
}
