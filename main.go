package main

import (
	"fmt"
	"log"

	accounts "github.com/bankAccount/banking"
)

func main() {
	// account := banking.Account{Owner: "dongunyun"} //balace는 조작할수있어 private
	// account.Owner = "pepe"// 변경하면됨
	//account := accounts.NewAccount("dongun")
	account := accounts.NewAccount("dongun")
	fmt.Println(*account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(10)
	//err := account.Withdraw(20)
	if err != nil { // 에러나오면 프로그램 종료
		log.Fatalln(err) // error에 내용을 Println을 해줌
	}
	fmt.Println(account.Balance(), account.Owner(), account.String())

}
