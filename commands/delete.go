package commands

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/models"
	"task-tracker/storage"
)

func Delete(args []string) {
	if len(args) == 0 {
		fmt.Println("Ошибка: укажите ID задачи для удаления.")
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

	var newTasks []models.Task
	found := false
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Ошибка: задача с ID %d не найдена.\n", id)
		os.Exit(1)
	}

	if err := storage.SaveTasks(newTasks); err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Задача удалена успешно.")
}
