package collections

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/clombo/Tutorials/TodoList/internal/db"
)

func Add(dbcon *sql.DB) {

	collectionName := promptName()

	err := db.AddCollection(dbcon, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Collection successfully added: %v\n", collectionName)
	}

}

func ShowTasks(dbcon *sql.DB) {

	collectionName := promptName()

	data, err := db.GetAllTasksByCollection(dbcon, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All tasks for collection: ", collectionName)
		for _, task := range *data {
			fmt.Printf("Number: %d, Task: %s\n", task.TaskNumber, task.Task)
		}
	}
}

func Remove(dbcon *sql.DB) {

	collectionName := promptName()
	err := db.RemoveCollection(dbcon, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Collection removed: ", collectionName)
	}
}

func promptName() string {
	//Prompt for collection name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter collection name: ")
	collectionName, _ := reader.ReadString('\n')
	collectionName = strings.TrimSpace(collectionName)

	return collectionName
}
