package routes

import (
	"kanban-backend/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/tasks", controllers.GetTasks)
	e.POST("/tasks", controllers.CreateTask)
	e.PUT("/tasks/:id", controllers.UpdateTask)
	e.PUT("/tasks/:id/deadline", controllers.UpdateDeadline) // ✅ Tambahkan update deadline
	e.DELETE("/tasks/:id", controllers.DeleteTask)
}
