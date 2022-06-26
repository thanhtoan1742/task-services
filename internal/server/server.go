package server

import (
	"context"

	"github.com/thanhtoan1742/task-services/api"
	"github.com/thanhtoan1742/task-services/internal/storage"
)

type Server struct {
	api.TaskServiceServer
}

func (s *Server) CheckHealth(ctx context.Context, r *api.EmptyRequest) (*api.ServiceHealth, error) {
	return &api.ServiceHealth{
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
