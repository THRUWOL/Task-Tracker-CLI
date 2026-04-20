package commands

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/models"
	"task-tracker/storage"
	"time"
)

func MarkInProgress(args []string) {
	markTask(args, models.StatusInProgress)
}

func MarkDone(args []string) {
	markTask(args, models.StatusDone)
}

func markTask(args []string, newStatus string) {
	if len(args) == 0 {
		fmt.Printf("Ошибка: укажите ID задачи.\n")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil || id <= 0 {
		fmt.Println("Ошибка: некорректный ID.")
		os.Exit(1)
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
		os.Exit(1)
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Ошибка: задача с ID %d не найдена.\n", id)
		os.Exit(1)
	}

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
		os.Exit(1)
	}

	statusRu := "в процессе"
	if newStatus == models.StatusDone {
		statusRu = "выполнена"
	}
	fmt.Printf("Статус задачи изменен на '%s'.\n", statusRu)
}
