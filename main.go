package main // main -> only for compiling

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
}
