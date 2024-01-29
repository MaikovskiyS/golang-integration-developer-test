package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"integration.v1/pkg/rpc"

	"integration.v1/gen/gamer"
	"integration.v1/internal/config"
	"integration.v1/internal/integration.v1/ports/freetogame"
	"integration.v1/internal/integration.v1/ports/repository"
	"integration.v1/internal/integration.v1/service"
	"integration.v1/internal/integration.v1/transport"
	"integration.v1/pkg/logger"
)

func Run(cfg *config.Config) error {
	logger, err := logger.New()
	if err != nil {
		return err
	}
	s := rpc.NewServer(logger)

	//init service deps
	freeToGameCl := freetogame.NewFreeToGameClient()
	repository := repository.NewMemoryStorage()
	svc := service.New(freeToGameCl, repository)

	// Register gRPC server
	gamer.RegisterServiceServer(s, transport.New(svc))

	go func() {
		fmt.Printf("gRPC server starting on :%s", cfg.Port)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
		if err != nil {
			logger.Error(err.Error())
		}

		if err = s.Serve(listener); err != nil {

			return
		}
	}()

	// shutdown.
	idleConnsClosed := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh

		s.GracefulStop()
		close(idleConnsClosed)
	}()

	<-idleConnsClosed
	return nil
}
