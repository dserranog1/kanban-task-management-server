package main

import (
	"github.com/dserranog1/kanban-task-management-server/db"
	"github.com/dserranog1/kanban-task-management-server/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := db.Init()
	router.InitBoardRoutes(e, db)
	e.Logger.Fatal(e.Start(":1323"))
}
