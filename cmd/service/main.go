package main

import (
	"fmt"
	"log"
	"skillfactory/postgresDB/storage"
)

var ConnString = "postgres://postgres:user@localhost:5432/tasks"

var id int
var rows []storage.Task

var task1 = storage.Task{
	AuthorId:   0,
	AssignedId: 0,
	Title:      "Some title",
	Content:    "Some content",
}

func main() {
	var db storage.InterfaceDB
	var err error

	db, err = storage.NewDb(ConnString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successful")

	id, err = db.NewTask(task1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	rows, err = db.AllTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(rows))
	for _, val := range rows {
		fmt.Println(val)
	}

	rows, err = db.TasksAuthor(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rows {
		fmt.Println(r)
	}

	rows, err = db.TasksLabel("label 2")
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range rows {
		fmt.Println(r)
	}

	err = db.UpdateTask(1, "new title", "new content")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task updated")

	err = db.RmTask(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task deleted")
}
