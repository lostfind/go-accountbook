package model

// MAccount is m_account model
type MAccount struct {
	AccountName string
	Balance     int
}

// SetBalance is Set Balance
func (m *MAccount) SetBalance(balance int) {
	m.Balance = balance
}
