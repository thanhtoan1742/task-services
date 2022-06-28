package server

import (
	"context"
	"io"
	"log"

	"github.com/thanhtoan1742/task-services/api"
	"github.com/thanhtoan1742/task-services/internal/storage"
)

type Server struct {
	api.TaskServiceServer
}

func (s *Server) CheckHealth(ctx context.Context, r *api.EmptyRequest) (*api.Status, error) {
	return &api.Status{
		Status:      0,
		Description: "Server is healthy.",
	}, nil
}

func (s *Server) GetTasks(ctx context.Context, r *api.EmptyRequest) (*api.TaskList, error) {
	tasks, err := storage.Read()
	if err != nil {
		return nil, err
	}

	tps := []*api.Task{}
	for i := 0; i < len(tasks); i++ {
		tps = append(tps, &tasks[i])
	}

	return &api.TaskList{
		Tasks: tps,
	}, nil
}

func (s *Server) GetTaskStream(r *api.EmptyRequest, stream api.TaskService_GetTaskStreamServer) error {
	log.Println(r)
	log.Println(stream)
	log.Println("Got stream request")
	tasks, err := storage.Read()
	if err != nil {
		return err
	}

	for i := 0; i < len(tasks); i++ {
		if err := stream.Send(&tasks[i]); err != nil {
			return err
		}
	}
	log.Println("Finished streaming")
	return nil
}

func (s *Server) AddTask(ctx context.Context, task *api.Task) (*api.Status, error) {
	log.Println("Got add task request")
	err := storage.Append(*task)
	if err != nil {
		return nil, err
	}
	log.Println("Added task")
	return &api.Status{
		Status:      0,
		Description: "ok",
	}, nil
}

func (s *Server) AddTaskStream(stream api.TaskService_AddTaskStreamServer) error {
	for {
		task, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := storage.Append(*task); err != nil {
			return err
		}

		status := api.Status{
			Status:      0,
			Description: "ok",
		}
		if err := stream.Send(&status); err != nil {
			return err
		}
	}
}
