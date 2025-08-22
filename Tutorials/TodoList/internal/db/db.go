package db

import (
	"database/sql"
	"fmt"

	"github.com/clombo/Tutorials/TodoList/internal/entities"
)

func Init(dbFile string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open DB connection")
	}

	// Ensure db is closed if we return early due to an error
	// This will be ignored if we successfully return the db at the end
	defer func() {
		if err != nil {
			db.Close()
		}
	}()

	// Create tasks table if it doesn’t exist
	createTasksTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		collection TEXT,
		task TEXT,
		taskNumber INTEGER
	);`

	_, err = db.Exec(createTasksTableSQL)
	if err != nil {
		return nil, fmt.Errorf("Failed to create tasks table: %v", err)
	}

	// Create collections table if it doesn't exist
	createCollectionsTableSQL := `CREATE TABLE IF NOT EXISTS collections(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	)`

	_, err = db.Exec(createCollectionsTableSQL)
	if err != nil {
		return nil, fmt.Errorf("Failed to create collections table: %v", err)
	}

	// If we reach here, everything succeeded, so don't close the db
	return db, nil
}

func GetTaskByID(db *sql.DB, id int) (*entities.Task, error) {

	query := "SELECT id, collection, task FROM tasks WHERE id = ?"

	var task entities.Task
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Collection, &task.Task)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task with ID %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving task: %v", err)
	}

	return &task, nil
}

func GetAllTasks(db *sql.DB) (*[]entities.Task, error) {

	query := "SELECT id, collection, task, taskNumber FROM tasks"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %v", err)
	}
	defer rows.Close() // Important: always close rows

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err := rows.Scan(&task.ID, &task.Collection, &task.Task, &task.TaskNumber)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %v", err)
		}
		tasks = append(tasks, task)
	}

	// Check for any error that occurred during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &tasks, nil
}

func GetAllTasksByCollection(db *sql.DB, collectionName string) (*[]entities.Task, error) {

	query := "SELECT id, collection, task, taskNumber FROM tasks WHERE collection = ?"

	rows, err := db.Query(query, collectionName)
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %v", err)
	}

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err := rows.Scan(&task.ID, &task.Task, &task.TaskNumber, &task.Collection)

		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %v", err)
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &tasks, nil

}

func CreateNewTask(db *sql.DB, description string) error {

}

func RemoveTask(db *sql.DB, id int) {

}

func GetCollections(db *sql.DB) {

}

func CollectionExists(db *sql.DB) {}

func AddCollection(db *sql.DB, collectionName string) {

}

func RemoveCollection(db *sql.DB, collectionName string) {

}

func generateUniqueId() int {
	//Generate a unique 4-digit ID starting from 1000
}
