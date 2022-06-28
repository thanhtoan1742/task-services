package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/thanhtoan1742/task-services/api"
	"github.com/thanhtoan1742/task-services/internal/client"
)

const (
	address string = "localhost"
	port    int32  = 10443
)

func checkHealth(c api.TaskServiceClient) {
	respond, err := c.CheckHealth(
		context.Background(),
		&api.EmptyRequest{},
	)
	if err != nil {
		log.Fatalf("Error when check server health: %v", err)
	}
	log.Printf("Got from server when check health: %s", respond.Description)
}

func getTask(c api.TaskServiceClient) {
	taskList, err := c.GetTasks(
		context.Background(),
		&api.EmptyRequest{},
	)
	if err != nil {
		log.Fatalf("Error when get tasks: %v", err)
	}
	log.Println("Got tasks from server:")
	for _, task := range taskList.Tasks {
		log.Printf("\t%v\n", task)
	}
}

func getTaskStream(c api.TaskServiceClient) {
	stream, err := c.GetTaskStream(context.Background(), &api.EmptyRequest{})
	if err != nil {
		log.Fatalf("Error when call get task stream: %v", err)
	}

	log.Println("Start receiving")
	for {
		task, err := stream.Recv()
		if err == io.EOF {
			log.Println("Received everything")
			break
		}
		if err != nil {
			log.Fatalf("Error when receiving: %v", err)
		}
		log.Println(task)
	}
}

func addTask(c api.TaskServiceClient) {
	status, err := c.AddTask(
		context.Background(),
		&api.Task{
			Name: "1111111",
		},
	)
	if err != nil {
		log.Fatalf("Error when add task: %v", err)
	}
	log.Println(status)
}

func addTaskStream(c api.TaskServiceClient) {
	stream, err := c.AddTaskStream(context.Background())
	if err != nil {
		log.Fatalf("Error when establish stream: %v\n", err)
	}
	for i := 0; i < 10; i++ {
		task := api.Task{
			Name: fmt.Sprintf("added task %d", i),
		}
		if err := stream.Send(&task); err != nil {
			log.Fatalf("Error when send task: %v\n", err)
		}
		status, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receive status: %v\n", err)
		}
		log.Println(status)
	}
}

func main() {
	log.Println("Started client")
	c := client.NewTaskServiceClient(address, port)
	// checkHealth(c)
	// getTask(c)
	addTask(c)
	// addTaskStream(c)
	getTaskStream(c)
}
