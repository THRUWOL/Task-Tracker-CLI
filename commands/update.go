package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker/storage"
	"time"
)

func Update(args []string) {
	if len(args) < 2 {
		fmt.Println("Ошибка: укажите ID и новое описание.")
		fmt.Println("Использование: task-cli update 1 \"Новое описание\"")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil || id <= 0 {
		fmt.Println("Ошибка: некорректный ID.")
		os.Exit(1)
	}

	newDesc := strings.Join(args[1:], " ")

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
		os.Exit(1)
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDesc
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

	fmt.Println("Задача обновлена успешно.")
}
