package accounts

type  Account struct {
	owner string
	balance int
}

// NewAccount 
func NewAccount(owner string) *Account {
	account := Account {owner:owner, balance:0}
	return &account
}