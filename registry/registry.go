package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/lostfind/go-accountbook/infrastructure/router"
)

type Interactor interface {
	NewHandlers() router.Handlers
}

type interactor struct {
	conn *gorm.DB
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{
		conn: conn,
	}
}

func (i *interactor) NewHandlers() router.Handlers {
	return router.Handlers{}
}
