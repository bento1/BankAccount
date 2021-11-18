package accounts

// Account Struct
type account struct { //대문자로 해야 import할수있다. // 근데 접근하면 안됨 함수로 접근하게
	owner   string //소문자면 private
	balance int
	// Owner   string //대문자면 public
	// Balance int
}

//NewAccount create Account
func NewAccount(owner string) *account { //2. 받은 애는 *account  객체 value로 받음
	account := account{owner: owner, balance: 0}
	return &account // 1.변수를 받은 넘이 새로운것을 복사하게 두기 싫음 주소로 준다.
}
