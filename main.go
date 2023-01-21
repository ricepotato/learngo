package main

import (
	"fmt"

	"github.com/ricepotato/learngo/accounts"

	"github.com/ricepotato/learngo/mydict"
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

	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	dictionary["first"] = "First Word"

	fmt.Println(dictionary)

	value, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value2, err2 := dictionary.Search("hello2")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(value2)
	}

	err3 := dictionary.Add("hello2", "hello2")
	if err3 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

	err4 := dictionary.Add("hello2", "hello2")
	if err4 != nil {
		fmt.Println("world already exists")
	} else {
		fmt.Println("success")
	}
}
