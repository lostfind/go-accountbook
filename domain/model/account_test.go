package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBalance(t *testing.T) {
	ma := new(Account)
	ma.SetBalance(100)

	assert.Equal(t, 100, ma.Balance)
}

func TestAccount_SetBalance(t *testing.T) {
	type fields struct {
		ID          int
		Name        string
		AccountType int
		Sort        int
		Balance     int
		DeleteFlag  bool
	}
	type args struct {
		balance int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "정상패턴 잔액 100 설정", args: args{balance: 100}, want: 100},
		{name: "정상패턴 잔액 0 설정", args: args{balance: 0}, want: 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &Account{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				AccountType: tt.fields.AccountType,
				Sort:        tt.fields.Sort,
				Balance:     tt.fields.Balance,
				DeleteFlag:  tt.fields.DeleteFlag,
			}
			m.SetBalance(tt.args.balance)
			if got := m.Balance; got != tt.want {
				t.Errorf("Account.Balance = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetTypeName(t *testing.T) {
	type fields struct {
		ID          int
		Name        string
		AccountType int
		Sort        int
		Balance     int
		DeleteFlag  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "자산구분 1 (Cash)", fields: fields{AccountType: Cash}, want: "Cash"},
		{name: "자산구분 2 (Bank)", fields: fields{AccountType: Bank}, want: "Bank"},
		{name: "자산구분 3 (Debit)", fields: fields{AccountType: Debit}, want: "Debit Card"},
		{name: "자산구분 4 (Credit)", fields: fields{AccountType: Credit}, want: "Credit Card"},
		{name: "자산구분 5 (E-Money)", fields: fields{AccountType: EMoney}, want: "Electric Money"},
		{name: "자산구분없음", fields: fields{AccountType: 99}, want: ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &Account{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				AccountType: tt.fields.AccountType,
				Sort:        tt.fields.Sort,
				Balance:     tt.fields.Balance,
				DeleteFlag:  tt.fields.DeleteFlag,
			}
			if got := m.GetTypeName(); got != tt.want {
				t.Errorf("Account.GetTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_UpdateBalance(t *testing.T) {
	type fields struct {
		ID          int
		Name        string
		AccountType int
		Sort        int
		Balance     int
		DeleteFlag  bool
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "잔액 감소", fields: fields{Balance: 50000}, args: args{-500}, want: 49500},
		{name: "잔액 증가", fields: fields{Balance: 50000}, args: args{8000}, want: 58000},
		{name: "변경 없음", fields: fields{Balance: 50000}, args: args{0}, want: 50000},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &Account{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				AccountType: tt.fields.AccountType,
				Sort:        tt.fields.Sort,
				Balance:     tt.fields.Balance,
				DeleteFlag:  tt.fields.DeleteFlag,
			}
			m.UpdateBalance(tt.args.amount)

			if got := m.Balance; got != tt.want {
				t.Errorf("Account.UpdateBalance() -> %v, want %v", got, tt.want)
			}
		})
	}
}
