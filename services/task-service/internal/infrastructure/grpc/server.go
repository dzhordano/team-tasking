package grpc

import (
	"log"
	"log/slog"
	"net"

	"github.com/dzhordano/team-tasking/services/tasks/internal/infrastructure/grpc/handlers"
	task_v1 "github.com/dzhordano/team-tasking/services/tasks/pkg/grpc/task/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	s    *grpc.Server
	port string
}

func NewServer(log *slog.Logger, ph *handlers.ProjectHandler, th *handlers.TaskHandler, ch *handlers.CommentHandler, port string, publickey []byte) *server {
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			//logging.StartCall, logging.FinishCall,
			logging.PayloadReceived, logging.PayloadSent,
		),
		// Add any other option (check functions starting with logging.With).
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...)),
		grpc.UnaryInterceptor(AuthInterceptor(publickey)))

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
