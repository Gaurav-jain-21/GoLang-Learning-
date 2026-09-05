package main

import "fmt"

type BankAccount struct{
	AccountNumber string 
	Owner string 
	Balance float64
}
func (b *BankAccount) Deposit(amount float64){
	if amount <=0{
		fmt.Println("Invalid amount")
		return 
	}
	b.Balance+=amount
	fmt.Printf("✅ Deposited ₹%.2f. New balance: ₹%.2f\n", amount, b.Balance)
}
func (b *BankAccount) Withdraw (amount float64){
	if amount<=0{
		fmt.Println("Invalid amount")
		return 
	}
	if amount> b.Balance{
		fmt.Println("Insufficent balance")
		return 
	}
	b.Balance -= amount
	fmt.Printf("✅ Withdrew ₹%.2f. New balance: ₹%.2f\n", amount, b.Balance)
}

func (b BankAccount) Display(){
	fmt.Printf("\n Account: %s\nOwner: %s\nBalance: %.2f\n",b.AccountNumber,b.Owner,b.Balance)
}

func main(){
	account:= BankAccount{
		AccountNumber: "ACC123456",
        Owner:         "Gaurav Sharma",
        Balance:       5000,
	}
	account.Display()
    account.Deposit(2500)
    account.Withdraw(1000)
    account.Withdraw(10000)  // Should fail
    account.Display()
}