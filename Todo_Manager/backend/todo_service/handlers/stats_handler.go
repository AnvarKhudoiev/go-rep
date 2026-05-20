package handlers

import (
	"backend/todo_service/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DayStats struct {
	Date      string `json:"date"`
	Label     string `json:"label"`
	Created   int64  `json:"created"`
	Completed int64  `json:"completed"`
}

type CategoryStats struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Count int64  `json:"count"`
}

type StatsResponse struct {
	Total          int64           `json:"total"`
	Completed      int64           `json:"completed"`
	Active         int64           `json:"active"`
	Overdue        int64           `json:"overdue"`
	CompletionRate int             `json:"completion_rate"`
	Week           []DayStats      `json:"week"`
	ByCategory     []CategoryStats `json:"by_category"`
}

func GetStats(c *gin.Context) {
	userID := c.GetUint("user_id")

	var total, completed, overdue int64

	config.DB.Table("todos").
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Count(&total)

	config.DB.Table("todos").
		Where("user_id = ? AND completed = true AND deleted_at IS NULL", userID).
		Count(&completed)

	config.DB.Table("todos").
		Where("user_id = ? AND completed = false AND due_date < ? AND due_date IS NOT NULL AND deleted_at IS NULL", userID, time.Now()).
		Count(&overdue)

	active := total - completed
	completionRate := 0
	if total > 0 {
		completionRate = int(completed * 100 / total)
	}

	weekdays := []string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"}
	week := make([]DayStats, 7)

	for i := 6; i >= 0; i-- {
		d := time.Now().AddDate(0, 0, -i)
		dateStr := d.Format("2006-01-02")
		label := weekdays[d.Weekday()]

		var created, comp int64
		config.DB.Table("todos").
			Where("user_id = ? AND DATE(created_at) = ? AND deleted_at IS NULL", userID, dateStr).
			Count(&created)
		config.DB.Table("todos").
			Where("user_id = ? AND completed = true AND DATE(updated_at) = ? AND deleted_at IS NULL", userID, dateStr).
			Count(&comp)

		week[6-i] = DayStats{Date: dateStr, Label: label, Created: created, Completed: comp}
	}

	type catRow struct {
		Name  string
		Color string
		Count int64
	}
	var catRows []catRow
	config.DB.Table("todos").
		Select("categories.name, categories.color, COUNT(todos.id) as count").
		Joins("LEFT JOIN categories ON todos.category_id = categories.id").
		Where("todos.user_id = ? AND todos.deleted_at IS NULL", userID).
		Group("categories.name, categories.color").
		Scan(&catRows)

	byCategory := make([]CategoryStats, 0, len(catRows))
	for _, row := range catRows {
		name := row.Name
		if name == "" {
			name = "Без категории"
		}
		color := row.Color
		if color == "" {
			color = "#94a3b8"
		}
		byCategory = append(byCategory, CategoryStats{Name: name, Color: color, Count: row.Count})
	}

	c.JSON(http.StatusOK, StatsResponse{
		Total:          total,
		Completed:      completed,
		Active:         active,
		Overdue:        overdue,
		CompletionRate: completionRate,
		Week:           week,
		ByCategory:     byCategory,
	})
}