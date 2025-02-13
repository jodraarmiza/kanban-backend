package main

import (
	"kanban-backend/config"
	"kanban-backend/models"
	"kanban-backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Task{}) // ✅ Pastikan model sudah di-migrate

	e := echo.New()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
