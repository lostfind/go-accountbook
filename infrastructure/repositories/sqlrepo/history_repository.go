package sqlrepo

import (
	"go-accountbook/domain/model"

	log "github.com/sirupsen/logrus"
)

func (r *sqlRepo) Find(id int) (*model.History, error) {
	history := &model.History{}

	selectSQL := `SELECT id, account_id, category_id, amount, memo, date
	FROM histories
	WHERE id = ?`

	stmt, err := r.db.Prepare(selectSQL)
	if err != nil {
		log.Error("HISTORY_REPO_FIND", err)
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

func (r *sqlRepo) FindAll() (model.Histories, error) {
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

func (r *sqlRepo) Save(history *model.History) (err error) {
	ins, err := r.db.Prepare("INSERT INTO histories(account_id,amount,memo) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	ins.Exec(history.AccountID, history.Amount, history.Memo)

	return nil
}

func (r *sqlRepo) FindAccounts(id []int) (map[int]*model.Account, error) {
	//TODO: implements
	account := make(map[int]*model.Account, len(id))

	for _, i := range id {
		account[i] = &model.Account{ID: i, Balance: 60000}
	}

	return account, nil
}
