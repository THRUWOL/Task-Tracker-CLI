package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker/models"
)

const dbFile = "tasks.json"

// LoadTasks загружает задачи из JSON файла
func LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Если файла нет, возвращаем пустой список — это не ошибка
			return []models.Task{}, nil
		}
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return tasks, nil
}

// SaveTasks сохраняет задачи в JSON файл
func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	// Атомарная запись: сначала пишем во временный файл, потом переименовываем
	tmpFile := dbFile + ".tmp"
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи временного файла: %w", err)
	}

	if err := os.Rename(tmpFile, dbFile); err != nil {
		return fmt.Errorf("ошибка переименования файла: %w", err)
	}

	return nil
}

// GetNextID вычисляет следующий свободный ID
func GetNextID(tasks []models.Task) int {
	if len(tasks) == 0 {
		return 1
	}
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
