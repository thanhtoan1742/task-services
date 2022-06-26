package server

import (
	"context"

	"github.com/thanhtoan1742/task-services/api"
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
	return &api.TaskList{
		Tasks: []*api.Task{
			{
				Name:        "Task 1",
				Description: "Description 1",
				DueTime:     "none",
				Finished:    false,
			},
		},
	}, nil
}
