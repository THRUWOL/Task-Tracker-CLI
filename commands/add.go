package commands

import (
	"fmt"
	"os"
	"strings"
	"task-tracker/models"
	"task-tracker/storage"
)

func Add(args []string) {
	if len(args) == 0 {
		fmt.Println("Ошибка: укажите описание задачи.")
		fmt.Println("Использование: task-cli add \"Купить молоко\"")
		os.Exit(1)
	}

	description := strings.Join(args, " ")

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
		os.Exit(1)
	}

	newTask := models.NewTask(description)
	newTask.ID = storage.GetNextID(tasks)

	tasks = append(tasks, newTask)

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Задача добавлена успешно (ID: %d)\n", newTask.ID)
}
