package model

// Account is m_account model
type Account struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"name"`
	AccountType int    `gorm:"account_type"`
	Sort        int    `gorm:"sort"`
	Balance     int    `gorm:"balance"`
	DeleteFlag  bool   `gorm:"is_delete"`
}

// NewAccount constructor Account
func NewAccount(name string, balance int, accountType int) *Account {
	return &Account{
		Name:        name,
		Balance:     balance,
		AccountType: accountType,
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
