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

	exists, err := db.CollectionExists(dbcon, collectionName)

	if err != nil {
		fmt.Println("Failed to add collection: ", err)
		return
	}

	if exists {
		fmt.Println("Collection already exists")
		return
	}

	err = db.AddCollection(dbcon, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Collection successfully added: %v\n", collectionName)
	}

}

func ShowTasks(dbcon *sql.DB) {

	collectionName := promptName()

	exists, err := db.CollectionExists(dbcon, collectionName)

	if err != nil {
		fmt.Println("Could not get tasks for collection: ", err)
		return
	}

	if !exists {
		fmt.Println("Collection does not exist")
		return
	}

	data, err := db.GetAllTasksByCollection(dbcon, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All tasks for collection: ", collectionName)
		fmt.Println("Task Number | Collection | Description")

		for _, task := range *data {
			fmt.Printf("%d | %s | %s\n", task.TaskNumber, task.Collection, task.Task)
		}
	}
}

func Remove(dbcon *sql.DB) {

	collectionName := promptName()

	exists, err := db.CollectionExists(dbcon, collectionName)

	if err != nil {
		fmt.Println("Error removing collection: ", err)
		return
	}

	if !exists {
		fmt.Println("Collection does not exist.")
		return
	}

	hasTasks, err := db.CollectionHasTasks(dbcon, collectionName)

	if err != nil {
		fmt.Println("Error removing collection: ", err)
		return
	}

	if hasTasks {
		fmt.Println("Please remove tasks first before deleting collection")
		return
	}

	err = db.RemoveCollection(dbcon, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Collection removed: ", collectionName)
	}
}

func DisplayCollections(dbcon *sql.DB) {

	collections, err := db.GetCollections(dbcon)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Available collections:")
		for _, collection := range *collections {
			fmt.Println(collection.Name)
		}
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
