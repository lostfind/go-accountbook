package model

import (
	"errors"
	"time"
)

type Budget struct {
	ID        int
	YearMonth time.Time
	Category  Category
	Amount    int
	Balance   int
}

func (b *Budget) SetYearMonth(ym string) error {
	layout := "2006-01"
	yearMonth, err := time.ParseInLocation(layout, ym, time.Local)

	if err == nil {
		b.YearMonth = yearMonth
		return nil
	}

	return errors.New("Parse err")
}
