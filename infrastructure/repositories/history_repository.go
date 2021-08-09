package repositories

import (
	"database/sql"
	"go-accountbook/domain/model"
	"go-accountbook/domain/repository"

	log "github.com/sirupsen/logrus"
)

type historyRepository struct {
	db *sql.DB
}

// NewHistoryRepository is constructor History Repository
func NewHistoryRepository(db *sql.DB) repository.HistoryRepository {
	return &historyRepository{
		db: db,
	}
}

func (r *historyRepository) Find(id int) (*model.History, error) {
	history := &model.History{}

	selectSQL := `SELECT id, account_id, category_id, amount, memo, date
	FROM histories
	WHERE id = ?`

	stmt, err := r.db.Prepare(selectSQL)
	if err != nil {
		log.Error(err)
		return history, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&history.ID, &history.AccountID, &history.CategoryID, &history.Amount, &history.Memo, &history.Date)
	if err != nil {
		log.Error(err)
		return history, err
	}

	return history, nil
}

func (r *historyRepository) FindAll() (model.Histories, error) {
	histories := model.Histories{}

	selectSQL := `SELECT id, account_id, category_id, amount, memo, date FROM histories`

	rows, err := r.db.Query(selectSQL)
	if err != nil {
		log.Error(err)
		return histories, err
	}
	defer rows.Close()
	for rows.Next() {
		history := &model.History{}
		rows.Scan(&history.ID, &history.AccountID, &history.CategoryID, &history.Amount, &history.Memo, &history.Date)

		histories = append(histories, history)
	}

	return histories, nil
}

func (r *historyRepository) Save(history *model.History) (err error) {
	return err
}
