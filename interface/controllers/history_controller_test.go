package controllers

import (
	"testing"
	"time"

	"github.com/lostfind/go-accountbook/conf"
	"github.com/lostfind/go-accountbook/entity/model"
	"github.com/lostfind/go-accountbook/infrastructure/datastore"
	"github.com/lostfind/go-accountbook/interface/repositories"
	"github.com/lostfind/go-accountbook/usecase/interactor"
	"github.com/stretchr/testify/assert"
)

func TestGetListHistory(t *testing.T) {
	conf.Read()
	db, _ := datastore.NewMySQLDB()
	repo := repositories.NewHistoryRepository(db)
	interactor := interactor.NewHistoryUsecase(repo)

	controller := NewHistoryController(interactor)

	histories, err := controller.GetListHistory()

	assert.NoError(t, err)
	assert.NotEmpty(t, histories)
}

func TestRegisterHistory(t *testing.T) {
	conf.Read()
	db, _ := datastore.NewMySQLDB()
	repo := repositories.NewHistoryRepository(db)
	interactor := interactor.NewHistoryUsecase(repo)

	controller := NewHistoryController(interactor)

	history := model.NewHistory(1, 1, 4000, "TEST", time.Now())
	err := controller.RegisterHistory(history)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, history.ID)
}
