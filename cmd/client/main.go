package main

import (
	"context"
	"log"

	"github.com/thanhtoan1742/task-services/api"
	"github.com/thanhtoan1742/task-services/internal/client"
)

const (
	address string = "localhost"
	port    int32  = 10443
)

func main() {
	log.Println("Started client")
	client := client.NewTaskServiceClient(address, port)

	respond, err := client.CheckHealth(
		context.Background(),
		&api.EmptyRequest{},
	)
	if err != nil {
		log.Fatalf("Error when check server health: %v", err)
	}
	log.Printf("Got from server when check health: %s", respond.Description)

	taskList, err := client.GetTasks(
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
