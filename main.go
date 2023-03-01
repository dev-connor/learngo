package main

import (
	"fmt"

	"github.com/dev-connor/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
