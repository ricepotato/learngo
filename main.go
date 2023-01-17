package main

import (
	"fmt"

	"github.com/ricepotato/learngo/accounts"
)

func main() {
	// account := banking.Account{Owner: "nico", Balance: 1000}
	//account.Owner = "pepe"
	//fmt.Println(account)
	account := accounts.NewAccount("ricepotato")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err) // just print error
		// log.Fatalln(err) // kill the program
	}
	fmt.Println(account.Balance())
	fmt.Println(account.Owner())

	account.ChangeOwner("new ricepotato")

	fmt.Println(account)

}
