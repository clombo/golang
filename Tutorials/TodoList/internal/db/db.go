package db

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/clombo/Tutorials/TodoList/internal/entities"
)

func Init(dbFile string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB connection")
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
		return nil, fmt.Errorf("failed to create tasks table: %v", err)
	}

	// Create collections table if it doesn't exist
	createCollectionsTableSQL := `CREATE TABLE IF NOT EXISTS collections(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	)`

	_, err = db.Exec(createCollectionsTableSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to create collections table: %v", err)
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

func CreateNewTask(db *sql.DB, description string, collection string) error {

	var numberExists bool
	numExistsQuery := "SELECT EXISTS(SELECT 1 FROM tasks WHERE taskNumber = ?)"
	var newTaskNumber int

	//Loop till you find number that does not exist on the database
	for {

		newTaskNumber = generateUniqueId()

		err := db.QueryRow(numExistsQuery, newTaskNumber).Scan(&numberExists)
		if err != nil {
			return fmt.Errorf("error creating new task: %v", err)
		}

		// If the task number doesn't exist, break out of loop
		if !numberExists {
			break
		}
	}

	// Insert the new task with the unique task number
	insertQuery := "INSERT INTO tasks (collection, task, taskNumber) VALUES (?, ?, ?)"
	_, err := db.Exec(insertQuery, collection, description, newTaskNumber)
	if err != nil {
		return fmt.Errorf("error inserting new task: %v", err)
	}

	return nil

}

func RemoveTask(db *sql.DB, id int) error {

	var exists bool
	existsQuery := "SELECT EXISTS(SELECT 1 FROM tasks WHERE id = ?)"

	err := db.QueryRow(existsQuery, id).Scan(&exists)

	if err != nil {
		return fmt.Errorf("error removing task: %v", err)
	}

	if !exists {
		return fmt.Errorf("task not found: %v", id)
	}

	// Delete the task
	deleteQuery := "DELETE FROM tasks WHERE id = ?"
	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting task: %v", err)
	}

	// Verify that exactly one row was affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task was deleted")
	}

	return nil
}

func GetCollections(db *sql.DB) (*[]entities.Collection, error) {

	query := "SELECT Name FROM collections"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting collections: %v", err)
	}

	var collections []entities.Collection

	for rows.Next() {
		var collection entities.Collection
		err := rows.Scan(&collection.Name)

		if err != nil {
			return nil, fmt.Errorf("error scanning collection row: %v", err)
		}

		collections = append(collections, collection)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &collections, nil
}

func CollectionExists(db *sql.DB, collection string) (bool, error) {

	var collectionExists bool
	query := "SELECT EXISTS(SELECT 1 FROM collections WHERE Name = ?)"

	err := db.QueryRow(query, collection).Scan(&collectionExists)

	if err != nil {
		return false, fmt.Errorf("error checking collection: %v", err)
	}

	return collectionExists, nil

}

func AddCollection(db *sql.DB, collectionName string) error {

	var collectionExists bool
	collectionName = strings.ToLower(collectionName)

	Insertquery := "INSERT INTO collections (name) VALUES (?)"
	existsQuery := "SELECT EXISTS(SELECT 1 FROM collections WHERE Name = ?)"

	err := db.QueryRow(existsQuery, collectionName).Scan(&collectionExists)

	if err != nil {
		return fmt.Errorf("error adding new collection: %v", err)
	}

	if !collectionExists {
		_, err := db.Exec(Insertquery, collectionName)

		if err != nil {
			return fmt.Errorf("error adding new collection: %v", err)
		}

	}

	return nil
}

func RemoveCollection(db *sql.DB, collectionName string) error {

	var collectionExists bool
	collectionName = strings.ToLower(collectionName)

	deleteQuery := "DELETE FROM collections WHERE name = ?"
	existsQuery := "SELECT EXISTS(SELECT 1 FROM collections WHERE Name = ?)"

	err := db.QueryRow(existsQuery, collectionName).Scan(&collectionExists)

	if err != nil {
		return fmt.Errorf("error removing collection: %v", err)
	}

	if collectionExists {
		_, err := db.Exec(deleteQuery, collectionName)

		if err != nil {
			return fmt.Errorf("error removing collection: %v", err)
		}
	}

	return nil

}

func generateUniqueId() int {
	return rand.IntN(9000) + 1000
}
