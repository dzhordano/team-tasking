package main

import (
	"log"

	services "github.com/dzhordano/team-tasking/services/tasks/internal/application/service"
	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc"
	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc/handlers"
)

func main() {
	project_service := services.NewProjectService(nil)
	task_service := services.NewTaskService(nil)
	comment_service := services.NewCommentService(nil)

	grpc_server := grpc.NewServer(
		handlers.NewProjectHandler(project_service),
		handlers.NewTaskHandler(task_service),
		handlers.NewCommentService(comment_service),
		"50000",
	)

	if err := grpc_server.Run(); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}
}
