package controllers

import (
	"fmt"
	"go-accountbook/domain/model"
	"go-accountbook/domain/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type historyController struct {
	usecase usecase.HistoryUsecase
}

type HistoryController interface {
	GetHistory(*gin.Context)
	GetHistories(ctx *gin.Context)
	RegisterHistory(history *model.History) error
}

// NewHistoryController is Constructor History controller
func NewHistoryController(i usecase.HistoryUsecase) HistoryController {
	return &historyController{
		usecase: i,
	}
}

func (c *historyController) RegisterHistory(history *model.History) error {
	return c.usecase.RegisterHistory(history)
}

func (c *historyController) GetHistories(ctx *gin.Context) {
	histories, err := c.usecase.GetHistories()
	if err != nil {
		fmt.Errorf("GetHistoriesERR : %v", err)
	}
	ctx.JSON(http.StatusOK, histories)
}

func (c *historyController) GetHistory(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)
	history, err := c.usecase.GetHistory(intId)
	if err != nil {
		log.Printf("GetHistoryERR : %v", err)
	}

	ctx.JSON(http.StatusOK, history)
}
