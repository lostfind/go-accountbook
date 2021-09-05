package model

// Account 자산구분 모델
type Account struct {
	ID          int
	Name        string
	AccountType int
	Sort        int
	Balance     int
	DeleteFlag  bool
}

const (
	Cash = iota + 1
	Bank
	Debit
	Credit
	EMoney
)

// SetBalance 잔액 설정
func (m *Account) SetBalance(balance int) {
	m.Balance = balance
}

// GetTypeName return Account type name
func (m *Account) GetTypeName() string {
	switch m.AccountType {
	case Cash:
		return "Cash"
	case Bank:
		return "Bank"
	case Debit:
		return "Debit Card"
	case Credit:
		return "Credit Card"
	case EMoney:
		return "Electric Money"
	}

	return ""
}

func (m *Account) UpdateBalance(amount int) {
	m.Balance += amount
}
