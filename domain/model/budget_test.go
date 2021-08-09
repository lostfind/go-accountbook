package model

import (
	"testing"
	"time"
)

func TestBudget_SetYearMonth(t *testing.T) {
	type fields struct {
		ID        int
		YearMonth time.Time
		Category  Category
		Amount    int
		Balance   int
	}
	type args struct {
		ym string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "", args: args{"2020-04"}, want: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local), wantErr: false},
		{name: "", args: args{"202005"}, wantErr: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := &Budget{
				ID:        tt.fields.ID,
				YearMonth: tt.fields.YearMonth,
				Category:  tt.fields.Category,
				Amount:    tt.fields.Amount,
				Balance:   tt.fields.Balance,
			}
			if err := b.SetYearMonth(tt.args.ym); (err != nil) != tt.wantErr {
				t.Errorf("Budget.SetYearMonth() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := b.YearMonth; got != tt.want {
				t.Errorf("Budget.SetYearMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
