package model

// Account is m_account model
type Account struct {
	ID      int    `gorm:"primary_key"`
	Name    string `gorm:"name"`
	Balance int    `gorm:"balance"`
}

// NewAccount constructor Account
func NewAccount(name string, balance int) *Account {
	return &Account{
		Name:    name,
		Balance: balance,
	}
}

// TableName return Table Name
func (Account) TableName() string {
	return "m_accounts"
}

// SetBalance is Set Balance
func (m *Account) SetBalance(balance int) {
	m.Balance = balance
}
