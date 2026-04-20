package main

import (
	"fmt"
	"os"
	"task-tracker/commands"
)

func printHelp() {
	var helpInfo = `----------------------------------------------------------
| <command> <args>      | description                     |
-----------------------------------------------------------
| add <description>     | Adding a new task               |
| update <id>           | Updating tasks by id            |
| delete <id>           | Delete tasks by id              |
| mark-in-progress <id> | Marking a task as in progress   |
| mark-done <id>        | Marking a task as done          |
| list             		| Listing all tasks               |
| list <status>   		| Listing task by status.         |
|                 		| Status: done, todo, in-progress |
-----------------------------------------------------------`
	fmt.Println(helpInfo)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		commands.Add(args)
	case "update":
		commands.Update(args)
	case "delete":
		commands.Delete(args)
	case "mark-in-progress":
		commands.MarkInProgress(args)
	case "mark-done":
		commands.MarkDone(args)
	case "list":
		commands.List(args)
	case "help":
		printHelp()
	default:
		fmt.Printf("Неизвестная команда: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}
