package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/lostfind/go-accountbook/entity/model"
	"github.com/lostfind/go-accountbook/entity/repository"
)

type historyRepository struct {
	db *gorm.DB
}

// NewHistoryRepository is constructor History Repository
func NewHistoryRepository(db *gorm.DB) repository.HistoryRepository {
	return &historyRepository{
		db: db,
	}
}

func (r *historyRepository) Find(id int) (history *model.History, err error) {
	err = r.db.First(&history, id).Error
	return
}

func (r *historyRepository) FindAll() (histories []*model.History, err error) {
	err = r.db.Find(&histories).Error
	return
}

func (r *historyRepository) Save(history *model.History) (err error) {
	err = r.db.Save(history).Error
	return err
}
