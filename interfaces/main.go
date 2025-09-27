package main

import "fmt"

type Account interface {
	GetBalance() int
	Deposit(amount int)
	WithDraw(amount int)
}

func main() {
	wf := NewWellFargo()
	var acct Account = wf // < ---  here Account is holding t evalue of the WellFargo means  it checks if the wf has all the methods,variables which are in the account interface
	acct.Deposit(1000)
	balance := acct.GetBalance()
	fmt.Println("The balance is the ", balance)
	acct.WithDraw(200)
	balance = acct.GetBalance()
	fmt.Println("The balance after withdraw is the ", balance)
}
