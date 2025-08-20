package db

import (
	"database/sql"

	"github.com/clombo/Tutorials/TodoList/internal/entities"
)

func Init(dbFile string) (*sql.DB, error) {

}

func GetTaskByID(db *sql.DB, id int) (*entities.Task, error) {

}

func GetAllTasks(db *sql.DB) (*[]entities.Task, error) {

}

func GetAllTasksByCollection(db *sql.DB, collectionName string) {

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
