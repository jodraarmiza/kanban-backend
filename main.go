package main

import (
	"fmt"
	"os"

	"kanban-backend/config"
	"kanban-backend/models"
	"kanban-backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Task{}) // Migrasi database otomatis

	e := echo.New()
	routes.InitRoutes(e)

	// Gunakan PORT dari environment atau default ke 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server berjalan di port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
