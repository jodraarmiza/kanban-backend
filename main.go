package main

import (
	"fmt"
	"os"

	"kanban-backend/config"
	"kanban-backend/models"
	"kanban-backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// ✅ Migrasikan model `User` dan `Task`
	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	e := echo.New()

	// ✅ Tambahkan Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://kanbantodolist.netlify.app"},                                                   // Ganti dengan domain frontend
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},                               // Tambahkan OPTIONS
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // Tambahkan Authorization jika ada login
	}))

	// Inisialisasi route
	routes.InitRoutes(e)

	// Gunakan PORT dari environment atau default ke 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server berjalan di port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
