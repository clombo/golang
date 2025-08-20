package tasks

import "fmt"

func Add() {
	fmt.Println("Adding a new task...")
}

func Remove(taskID int) {
	fmt.Printf("Removing task with ID: %d\n", taskID)
}

func ShowAll() {
	fmt.Println("Showing all tasks...")
}
