package registry

import (
	"database/sql"
	"go-accountbook/infrastructure/router"
)

type Interactor interface {
	NewHandlers() router.Handlers
}

type interactor struct {
	conn *sql.DB
}

func NewInteractor(conn *sql.DB) Interactor {
	return &interactor{
		conn: conn,
	}
}

func (i *interactor) NewHandlers() router.Handlers {
	return router.Handlers{}
}
