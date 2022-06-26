package main

import (
	"fmt"

	"github.com/thanhtoan1742/task-services/api"
	"github.com/thanhtoan1742/task-services/internal/storage"
)

func printTask() {
	tasks, err := storage.Read()
	if err != nil {
		fmt.Println(err)
	}
	for _, task := range tasks {
		fmt.Printf("%+v\n", task)
	}
}

func writeTasks() {
	tasks := []api.Task{}
	for i := 0; i < 10; i++ {
		tasks = append(tasks, api.Task{
			Name:        fmt.Sprintf("task %d", i),
			Description: fmt.Sprintf("description %d", i),
			DueTime:     "none",
		})
	}
	storage.Write(tasks)
}

func main() {
	writeTasks()
	printTask()
}
