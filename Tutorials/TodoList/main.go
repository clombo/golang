package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clombo/Tutorials/TodoList/collections"
	"github.com/clombo/Tutorials/TodoList/internal/db"
	"github.com/clombo/Tutorials/TodoList/tasks"

	_ "github.com/mattn/go-sqlite3"
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

	//Initialize DB connection
	dbcon, err := db.Init("./tododb")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	mainMenu()

	for {

		fmt.Print("\nEnter a command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "add":
			tasks.Add(dbcon)
		case "add collection":
			collections.Add(dbcon)
		case "show":
			tasks.ShowAll(dbcon)
		case "show collection":
			collections.ShowTasks(dbcon)
		case "display":
			collections.DisplayCollections(dbcon)
		case "remove task":
			tasks.Remove(dbcon) // Example task ID, replace with actual logic to get task ID
		case "remove collection":
			collections.Remove(dbcon) // Assuming "default" is a placeholder for a specific collection
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
	fmt.Println("3. display - Display all available collections")
	fmt.Println("4. show - Display all tasks across all collections.")
	fmt.Println("5. show collection - Display all tasks within a specific collection.")
	fmt.Println("6. remove task - Remove a specific task by its ID.")
	fmt.Println("7. remove collection - Remove a collection and all its tasks. If the collection has tasks, you will be prompted to confirm the deletion.")
	fmt.Println("8. help - Display this help menu.")
	fmt.Println("9. exit - Exit the application.")
}
