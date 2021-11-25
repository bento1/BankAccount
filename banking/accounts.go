package accounts

import "errors"

// Account Struct
type account struct { //대문자로 해야 import할수있다. // 근데 접근하면 안됨 함수로 접근하게
	owner   string //소문자면 private
	balance int
	// Owner   string //대문자면 public
	// Balance int
}

var errNoMoney = errors.New("cant withdraw you are poor")

//NewAccount create Account
func NewAccount(owner string) *account { //2. 받은 애는 *account  객체 value로 받음
	account := account{owner: owner, balance: 0}
	return &account // 1.변수를 받은 넘이 새로운것을 복사하게 두기 싫음 주소로 준다.
}

//method는 func와 좀 다름
//func 와 이름 사이에 리시버 () 를 달아준다.
//Deposit x amount on your account
//보통 이름은 struct의 첫글자의 소문자
// () 호출한 객체의 포인터가 가르키는 구조체를 사용해라
// 반영이 안되는 이유는 account만할떄 계쏙 복사해서 쓰기때문에.
//
func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a *account) Withdraw(amount int) error {
	if a.balance >= amount {
		a.balance -= amount
		return nil //에러 타입일떄  반대 출력 none 의미 꼭해줌
	} else {
		return errNoMoney
	}

}
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//읽기만 하는건 새로운 걸 복사하는게 맞는듯
func (a account) Owner() string {
	return a.owner
}
func (a account) Balance() int {
	return a.balance
}

func (a account) String() string {
	return (a.owner)
}
