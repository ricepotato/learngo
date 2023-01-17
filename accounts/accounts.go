package accounts

// Account struct private
type Account struct {
	owner   string // private
	balance int
	// Age int // public
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
