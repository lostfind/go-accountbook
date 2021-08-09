package repositories

import (
	"database/sql"
	"go-accountbook/domain/model"
	"go-accountbook/domain/repository"
	"reflect"
	"testing"
)

func TestNewHistoryRepository(t *testing.T) {
	testDB := InitTestDB()
	defer testDB.Close()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want repository.HistoryRepository
	}{
		{name: "정상패턴", args: args{testDB}, want: &historyRepository{testDB}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHistoryRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHistoryRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_historyRepository_Find(t *testing.T) {
	testDB := InitTestDB()
	defer testDB.Close()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantHistory *model.History
		wantErr     bool
	}{
		{name: "ID가 존재할 경우", fields: fields{testDB}, args: args{3}, wantHistory: &model.History{ID: 3, AccountID: 333, CategoryID: 3333, Amount: 1214, Memo: "테스트항목3", Date: "2021/08/08"}, wantErr: false},
		{name: "ID가 존재하지 않을 경우", fields: fields{testDB}, args: args{99}, wantHistory: &model.History{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &historyRepository{
				db: tt.fields.db,
			}
			gotHistory, err := r.Find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("historyRepository.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHistory, tt.wantHistory) {
				t.Errorf("historyRepository.Find() = %v, want %v", gotHistory, tt.wantHistory)
			}
		})
	}
}

func Test_historyRepository_FindAll(t *testing.T) {
	testDB := InitTestDB()
	defer testDB.Close()

	emptyDB := InitEmptyTestDB()
	defer emptyDB.Close()

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Histories
		wantErr bool
	}{
		{
			name:   "데이터가 있을 경우",
			fields: fields{testDB},
			want: model.Histories{
				&model.History{ID: 1, AccountID: 111, CategoryID: 1111, Amount: 50000, Memo: "테스트항목1", Date: "2021/08/09"},
				&model.History{ID: 2, AccountID: 222, CategoryID: 2222, Amount: 51230, Memo: "테스트항목2", Date: "2021/08/08"},
				&model.History{ID: 3, AccountID: 333, CategoryID: 3333, Amount: 1214, Memo: "테스트항목3", Date: "2021/08/08"},
			},
			wantErr: false,
		},
		{
			name:    "데이터가 없을 경우",
			fields:  fields{emptyDB},
			want:    model.Histories{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &historyRepository{
				db: tt.fields.db,
			}
			got, err := r.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("historyRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("historyRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
