package client

import (
	"fmt"
	"log"

	"github.com/thanhtoan1742/task-services/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTaskServiceClient(address string, port int32) api.TaskServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", address, port), opts...)
	if err != nil {
		log.Fatalf("Cannot create client because: %v\n", err)
	}
	return api.NewTaskServiceClient(conn)
}
