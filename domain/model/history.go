package model

// History 이력 모델
type History struct {
	ID         int    `json:"id,omitempty"`
	AccountID  int    `json:"account_id,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
	Amount     int    `json:"amount,omitempty"`
	Memo       string `json:"memo,omitempty"`
	Date       string `json:"date,omitempty"`
}

type Histories []*History
