package main 

import "fmt"

type BankAccount struct {
	Owner string 
	Balance float64
}
func (a BankAccount) GetBalance() float64{
	return a.Balance
}

func (a *BankAccount) Deposit(amount float64){
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64)string{
	if amount > a.Balance {
        return "❌ Insufficient funds!"
    }
    a.Balance -= amount
    return "✅ Withdrawal successful!"
}
func class28() {
    account := BankAccount{Owner: "Gaurav", Balance: 1000}
    
    fmt.Printf("Initial balance: ₹%.2f\n", account.GetBalance())
    
    account.Deposit(500)
    fmt.Printf("After deposit: ₹%.2f\n", account.GetBalance())
    
    msg := account.Withdraw(200)
    fmt.Println(msg)
    fmt.Printf("After withdrawal: ₹%.2f\n", account.GetBalance())
    
    msg2 := account.Withdraw(2000)
    fmt.Println(msg2)
}