package main // main -> only for compiling

import (
	"fmt"
	"learngo/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
