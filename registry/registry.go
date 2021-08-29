package registry

import (
	"go-accountbook/domain/usecase"
	"go-accountbook/infrastructure/datastore"
	"go-accountbook/infrastructure/repositories"
	"go-accountbook/interface/api/controllers"
)

func InjectionHistory() controllers.HistoryController {
	db := datastore.NewSqliteDB()
	repo := repositories.NewHistoryRepository(db)
	usecase := usecase.NewHistoryUsecase(repo)

	return controllers.NewHistoryController(usecase)
}
