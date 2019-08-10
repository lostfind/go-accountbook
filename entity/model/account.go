package model

// Account is m_account model
type Account struct {
	ID          int
	AccountName string
	Balance     int
}

// NewAccount constructor Account
func NewAccount(name string, balance int) *Account {
	return &Account{
		AccountName: name,
		Balance:     balance,
	}
}

// SetBalance is Set Balance
func (m *Account) SetBalance(balance int) {
	m.Balance = balance
}
