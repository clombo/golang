package tasks

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/clombo/Tutorials/TodoList/internal/db"
)

func Add(dbcon *sql.DB) {

	collectionName := promptForCollection()
	newTask := promptForTaskDetails()

	colExists, err := db.CollectionExists(dbcon, collectionName)

	if err != nil {
		fmt.Println("Failed to add task: ", err)
		return
	}

	if !colExists {
		fmt.Println("Cannot add task to collection that does not exist.")
		return
	}

	err = db.CreateNewTask(dbcon, newTask, collectionName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New task added successfully")
	}
}

func Remove(dbcon *sql.DB) {
	taskNumber, err := promptForTaskNumber()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.RemoveTask(dbcon, taskNumber)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Task sucessfully removed: ", taskNumber)
	}
}

func ShowAll(dbcon *sql.DB) {

	tasks, err := db.GetAllTasks(dbcon)

	if err != nil {
		fmt.Println("Failed getting all tasks: ", err)
		return
	}

	if len(*tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}

	fmt.Println("Task Number | Collection | Description")

	for _, task := range *tasks {
		fmt.Printf("%d | %s | %s\n", task.TaskNumber, task.Collection, task.Task)
	}

}

func promptForCollection() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter collection name: ")
	collectionName, _ := reader.ReadString('\n')
	collectionName = strings.TrimSpace(collectionName)

	return collectionName

}

func promptForTaskNumber() (int, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNumber, err := strconv.Atoi(input)

	if err != nil {
		return 0, fmt.Errorf("invalid number format. Please enter a valid task number")
	}

	return taskNumber, nil
}

func promptForTaskDetails() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task description: ")
	taskDetail, _ := reader.ReadString('\n')
	taskDetail = strings.TrimSpace(taskDetail)

	return taskDetail
}
