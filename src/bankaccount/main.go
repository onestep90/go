package main

import (
	"bankaccount/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("onestep")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
	account.ChangeOwner("other")
	fmt.Println(account)

	account2 := accounts.NewAccount("hihi")
	fmt.Println(account2)
	fmt.Println(account)
}
