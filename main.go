package main

import (
	"fmt"

	"github.com/moonssa/learngo/banking"
)




func main() {
	account := banking.Account{Owner:"moon", Balance:1000}
	fmt.Println(account)
}