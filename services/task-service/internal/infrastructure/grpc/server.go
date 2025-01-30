package grpc

import (
	"log"
	"net"

	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc/handlers"
	task_v1 "github.com/dzhordano/team-tasking/services/tasks/pkg/grpc/task/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	s    *grpc.Server
	port string
}

func NewServer(ph *handlers.ProjectHandler, th *handlers.TaskHandler, ch *handlers.CommentHandler, port string) *server {
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(s)

	task_v1.RegisterProjectServiceV1Server(s, ph)
	task_v1.RegisterTaskServiceV1Server(s, th)
	task_v1.RegisterCommentServiceV1Server(s, ch)

	return &server{
		s:    s,
		port: ":" + port,
	}
}

func (s *server) Run() error {
	list, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	log.Printf("starting gRPC server on port %s", s.port)

	return s.s.Serve(list)
}

func (s *server) Stop() {
	s.s.GracefulStop()
}
