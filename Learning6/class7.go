package main

import "fmt"
type Wallet struct{
	Balance float64
}

func (w *Wallet) Deposit(amount float64){
	w.Balance= w.Balance + amount
	fmt.Printf("Deposited $%.2f\n",amount)
}
func (w Wallet) GetBalance()float64{
	return w.Balance
}

func class7(){
	myWallet:=Wallet{
		Balance: 5000,
	}
	myWallet.Deposit(500)
	current:=myWallet.GetBalance()

	fmt.Printf("Current Balance $%.2f\n",current)
}