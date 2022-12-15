package router

import (
	"github.com/dserranog1/kanban-task-management-server/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitBoardRoutes(e *echo.Echo, db *gorm.DB) {
	boardHandler := handler.NewBoardHandler(db)
	e.GET("/boards", boardHandler.GetAllBoards)
}
