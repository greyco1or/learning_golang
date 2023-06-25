package main

import (
	"fmt"
	"log"

	"github.com/greyco1or/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("grey")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}