package controllers

import (
	"kanban-backend/config"
	"kanban-backend/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// ✅ Cek apakah username ada di database
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username atau password salah"})
	}

	// ✅ Cek password (pastikan menggunakan hash jika sudah terenkripsi)
	if user.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username atau password salah"})
	}

	// ✅ Jika login berhasil, kirim respon
	return c.JSON(http.StatusOK, map[string]string{"message": "Login berhasil!", "username": user.Username})
}

// ✅ 1. API untuk mengambil tugas yang masih dalam rentang createdAt hingga deadline
func GetTasks(c echo.Context) error {
	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	if startDate == "" || endDate == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Start and end date are required"})
	}

	var tasks []models.Task
	config.DB.Where("(created_at <= ? AND deadline >= ?)", endDate, startDate).Find(&tasks)

	return c.JSON(http.StatusOK, tasks)
}

// ✅ 2. API untuk membuat tugas & menyebarkannya ke semua tanggal dari createdAt hingga deadline
func CreateTask(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return err
	}

	// Konversi string ke time.Time
	startDate, _ := time.Parse("2006-01-02", task.CreatedAt)
	deadlineDate, _ := time.Parse("2006-01-02", task.Deadline)

	// Simpan tugas untuk setiap hari dari createdAt hingga deadline
	for !startDate.After(deadlineDate) {
		dateKey := startDate.Format("2006-01-02")

		newTask := *task
		newTask.CreatedAt = dateKey
		newTask.ID = 0 // Gunakan ID baru agar tidak menimpa task sebelumnya
		config.DB.Create(&newTask)

		startDate = startDate.AddDate(0, 0, 1) // Tambah 1 hari
	}

	return c.JSON(http.StatusCreated, task)
}

// ✅ 3. API untuk mengupdate tugas (termasuk status & deskripsi)
func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var task models.Task
	config.DB.First(&task, id)

	if task.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	c.Bind(&task)
	config.DB.Save(&task)
	return c.JSON(http.StatusOK, task)
}

// ✅ 4. API untuk memperbarui deadline & menyebarkan ulang tugas ke semua tanggal
func UpdateDeadline(c echo.Context) error {
	id := c.Param("id")
	var task models.Task
	config.DB.First(&task, id)

	if task.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	var requestData struct {
		Deadline string `json:"deadline"`
	}
	if err := c.Bind(&requestData); err != nil {
		return err
	}

	// Perbarui deadline di database
	task.Deadline = requestData.Deadline
	config.DB.Save(&task)

	// Tambahkan kembali tugas di setiap tanggal dari createdAt hingga deadline
	startDate, _ := time.Parse("2006-01-02", task.CreatedAt)
	deadlineDate, _ := time.Parse("2006-01-02", task.Deadline)

	for !startDate.After(deadlineDate) {
		dateKey := startDate.Format("2006-01-02")

		newTask := task
		newTask.CreatedAt = dateKey
		newTask.ID = 0 // Buat ID baru agar tidak menimpa task sebelumnya
		config.DB.Create(&newTask)

		startDate = startDate.AddDate(0, 0, 1) // Tambah 1 hari
	}

	return c.JSON(http.StatusOK, task)
}

// ✅ 5. API untuk menghapus tugas berdasarkan ID
func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(&models.Task{})
	return c.NoContent(http.StatusNoContent)
}
