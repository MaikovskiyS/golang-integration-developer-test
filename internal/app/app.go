package app

import (
	"fmt"
	"net"

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

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return err
	}

	fmt.Printf("gRPC server starting on :%s", cfg.Port)
	if err := s.Serve(l); err != nil {
		return err
	}
	return nil
}
