package main // main -> only for compiling

import (
	"fmt"
	"learngo/banking"
)

func main() {
	account := banking.Account{Owner: "sunbin", Balance: 1000}
	fmt.Println(account)
}
