package server

import (
	"fmt"
	"net"
	"os"

	"github.com/mikestefanello/surveys-microservices/survey-service/config"
	"github.com/mikestefanello/surveys-microservices/survey-service/handler"
	protos "github.com/mikestefanello/surveys-microservices/survey-service/protos/survey"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer starts a new gRPC server
func StartGrpcServer(h *handler.SurveyGrpcHandler, cfg config.GrpcConfig, log *zerolog.Logger) {
	gs := grpc.NewServer()
	protos.RegisterSurveyServer(gs, h)
	reflection.Register(gs)

	addr := fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port)
	l, err := net.Listen(cfg.Network, addr)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to start gRPC server.")
		os.Exit(1)
	}
	log.Info().Str("on", addr).Msg("Starting gRPC server")

	gs.Serve(l)
}
