package accounts

import "errors"

// Account struct private
type Account struct {
	owner   string // private
	balance int
	// Age int // public
}

var errNoMoney = errors.New("Can't withdraw you are poor.")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
