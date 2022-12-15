package handler

import (
	"net/http"

	"github.com/dserranog1/kanban-task-management-server/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BoardHandler struct {
	db *gorm.DB
}

// constructor
func NewBoardHandler(db *gorm.DB) *BoardHandler {
	return &BoardHandler{db}
}

// methods
func (boardHandler BoardHandler) GetAllBoards(c echo.Context) error {
	boardModel := model.NewBoardModel(boardHandler.db)
	boards := boardModel.GetAllBoards()
	return c.JSON(http.StatusOK, boards)
}
