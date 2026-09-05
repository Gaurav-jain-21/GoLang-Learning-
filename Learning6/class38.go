package main

import "fmt"

func withdraw(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, fmt.Errorf("invalid amount : %.2f (must be positive)", amount)
	}
	if amount> balance{
		return balance, fmt.Errorf("insufficent balance: have %.2f, want %.2f",balance  , amount)
	}
	return balance-amount, nil
}
func class38(){
	newBalance , err:=withdraw(1000,5000)
	if err!=nil{
		fmt.Println("Transcation failed",err)
		return 
	}
	fmt.Println("new balance : ",newBalance)
}