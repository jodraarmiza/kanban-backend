package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   string `json:"createdAt"`
	Deadline    string `json:"deadline"`
}
