package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clombo/Tutorials/TodoList/collections"
	"github.com/clombo/Tutorials/TodoList/tasks"
)

/*
Project planning:
Commands
- exit
- add "task" "collection"
- add "collection"
- show
- show "collection"
- remove "taskId"
- remove "collection"
- help

Each task will generate a unique 4 digit ID that starts at 1000.
Tasks will be stored in collections.

All data will be stored in a sqlite database.

Rules:
If collection has tasks and try and delete it, it will prompt the user to confirm deletion.
If collection is empty, it will be deleted without confirmation.
If task is removed, it will be removed without confirmation.
Task cannot be added without a collection and collection needs to exist.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	mainMenu()

	for {

		fmt.Print("\nEnter a command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "add":
			tasks.Add()
		case "add collection":
			collections.Add()
		case "show":
			tasks.ShowAll()
		case "show collection":
			collections.ShowTasks("default") // Assuming "default" is a placeholder for a specific collection
		case "remove task":
			tasks.Remove(1001) // Example task ID, replace with actual logic to get task ID
		case "remove collection":
			collections.Remove("default") // Assuming "default" is a placeholder for a specific collection
		case "help":
			mainMenu()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Type 'help' to see available commands.")
		}
	}
}

func mainMenu() {
	fmt.Println("### Welcome to the Todo List Application! ###")
	fmt.Println()
	fmt.Println("# Main Menu #")
	fmt.Println("1. add - Add a new task to a collection. This will require you to specify the task and the collection it belongs to.")
	fmt.Println("2. add collection - Create a new collection to organize your tasks.")
	fmt.Println("3. show - Display all tasks across all collections.")
	fmt.Println("4. show collection - Display all tasks within a specific collection.")
	fmt.Println("5. remove task - Remove a specific task by its ID.")
	fmt.Println("6. remove collection - Remove a collection and all its tasks. If the collection has tasks, you will be prompted to confirm the deletion.")
	fmt.Println("7. help - Display this help menu.")
	fmt.Println("8. exit - Exit the application.")
}
