package main

import (
	"fmt"

	accounts "github.com/bankAccount/banking"
)

func main() {
	// account := banking.Account{Owner: "dongunyun"} //balace는 조작할수있어 private
	// account.Owner = "pepe"// 변경하면됨
	//account := accounts.NewAccount("dongun")
	account := accounts.NewAccount("dongun")
	fmt.Println(*account)
}
