package model

// History 이력 모델
type History struct {
	ID         int
	AccountID  int
	CategoryID int
	Amount     int
	Memo       string
	Date       string
}

type Histories []*History
