package commands

import (
	"fmt"
	"os"
	"task-tracker/models"
	"task-tracker/storage"
)

func List(args []string) {
	filter := ""
	if len(args) > 0 {
		filter = args[0]
		if filter != models.StatusTodo &&
			filter != models.StatusInProgress &&
			filter != models.StatusDone {
			fmt.Println("Ошибка: неверный фильтр. Используйте: todo, in-progress, done")
			os.Exit(1)
		}
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	foundAny := false
	fmt.Printf("%-5s %-15s %-30s %-20s\n", "ID", "Статус", "Описание", "Обновлено")
	fmt.Println("---------------------------------------------------------------")

	for _, task := range tasks {
		if filter != "" && task.Status != filter {
			continue
		}

		foundAny = true
		fmt.Printf("%-5d %-15s %-30s %-20s\n",
			task.ID,
			task.Status,
			truncateString(task.Description, 30),
			models.FormatDateFriendly(task.UpdatedAt))
	}

	if !foundAny {
		fmt.Println("Нет задач с указанным статусом.")
	}
}

func truncateString(s string, maxlen int) string {
	if len(s) <= maxlen {
		return s
	}
	return s[:maxlen-3] + "..."
}
