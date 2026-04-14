package main

import (
	"fmt"
	"os"
)

var helpInfo = 
`------------------------------------------------------
| <command>        | description                     |
------------------------------------------------------
| add              | Adding a new task               |
| update           | Updating tasks                  |
| delete           | Delete tasks                    |
| mark-in-progress | Marking a task as in progress   |
| mark-done        | Marking a task as done          |
| list             | Listing all tasks               |
| list <status>    | Listing task by status.         |
|                  | Status: done, todo, in-progress |
------------------------------------------------------`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpInfo)
		return
	}

	command := os.Args[1]

	switch command {
		case "help", "man", "-h", "info", "-m":
			fmt.Println(helpInfo)
		case "add":
			//todo add
		case "update":
			//todo update
		case "delete":
			//todo delete
		case "mark":
			//todo
		case "list":
			//todo
		default:
			fmt.Println("Unknown command: ", command)
	}

}
