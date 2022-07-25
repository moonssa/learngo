package main

import (
	"fmt"

	"github.com/moonssa/learngo/accounts"
)




func main() {
	account := accounts.NewAccount("moon")
	
	fmt.Println(account)
}