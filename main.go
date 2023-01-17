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
}
