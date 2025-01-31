package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	services "github.com/dzhordano/team-tasking/services/tasks/internal/application/service"
	"github.com/dzhordano/team-tasking/services/tasks/internal/config"
	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc"
	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc/handlers"
	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/repository"
)

func main() {
	// Не забыть сделать генерацию таких ключей при деплое (volume в kubernetes)
	publickey, err := os.ReadFile("jwt.key.pub")
	if err != nil {
		log.Fatalf("failed to read public key: %v", err)
	}

	cfg := config.MustLoad()

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	pool, err := repository.NewPGXPool(context.Background(), cfg.PG.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	project_repo := repository.NewPGProjectRepository(pool)
	task_repo := repository.NewPGTaskRepository(pool)
	comment_repo := repository.NewPGCommentRepository(pool)

	project_service := services.NewProjectService(project_repo)
	task_service := services.NewTaskService(task_repo, project_repo)
	comment_service := services.NewCommentService(comment_repo)

	grpc_server := grpc.NewServer(
		logger,
		handlers.NewProjectHandler(project_service),
		handlers.NewTaskHandler(task_service),
		handlers.NewCommentService(comment_service),
		cfg.GRPC.Port,
		publickey,
	)

	if err := grpc_server.Run(); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}
}
