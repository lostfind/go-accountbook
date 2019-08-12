package model

// Account is m_account model
type Account struct {
	ID          int
	AccountName string
	AccountType int
	Sort        int
	Balance     int
	DeleteFlag  bool
}

// TableName return Table Name
func (Account) TableName() string {
	return "m_account"
}

// NewAccount is constructor Account
func NewAccount(name string, balance, accountType int) *Account {
	return &Account{
		AccountName: name,
		Balance:     balance,
		AccountType: accountType,
	}
}

// SetBalance is Set Balance
func (m *Account) SetBalance(balance int) {
	m.Balance = balance
}

// GetTypeName return Account type name
func (m *Account) GetTypeName() string {
	switch m.AccountType {
	case 1:
		return "Cash"
	case 2:
		return "Bank"
	case 3:
		return "Debit Card"
	case 4:
		return "Credit Card"
	case 5:
		return "Electric Money"
	}

	return ""
}
