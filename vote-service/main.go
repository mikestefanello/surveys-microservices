package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	protos "github.com/mikestefanello/surveys-microservices/survey-service/protos/survey"
	"google.golang.org/grpc"

	"github.com/mikestefanello/surveys-microservices/vote-service/config"
	"github.com/mikestefanello/surveys-microservices/vote-service/handler"
	"github.com/mikestefanello/surveys-microservices/vote-service/logger"
	"github.com/mikestefanello/surveys-microservices/vote-service/repository"
	"github.com/mikestefanello/surveys-microservices/vote-service/router"
	"github.com/mikestefanello/surveys-microservices/vote-service/serializer"
	"github.com/mikestefanello/surveys-microservices/vote-service/server"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
)

func main() {
	// Get a logger
	log := logger.NewConsoleLogger()

	// Load application configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load application configuration")
		os.Exit(1)
	}

	// Load the repository
	sz := serializer.NewVoteJSONSerializer()
	repo, err := repository.NewRabbitVoteWriterRepository(cfg.Rabbit, sz)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to repository")
		os.Exit(1)
	}

	// Connect to the survey gRPC service
	grpcAddr := fmt.Sprintf("%s:%d", cfg.SurveyGrpc.Hostname, cfg.SurveyGrpc.Port)
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to survey gRPC service")
		os.Exit(1)
	}
	cli := protos.NewSurveyClient(conn)
	log.Info().Str("on", grpcAddr).Msg("Connected to survey gRPC service")

	// Load the service
	service := vote.NewService(repo, repo, cli)

	// Load HTTP dependencies
	httpHandler := handler.NewVoteHTTPHandler(service, &log)
	httpRouter := router.NewRouter(httpHandler)
	httpServer := server.NewHTTPServer(httpRouter, cfg.HTTP)

	// Start the HTTP server
	go func() {
		log.Info().Str("on", httpServer.Addr).Msg("Starting HTTP server")
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal().Err(err).Msg("Server shutdown")
			os.Exit(1)
		}
	}()

	// Listen for sigterm or interupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received
	sig := <-c
	log.Warn().Msgf("Signal received: %v", sig)

	// Gracefully shutdown the server allowing up to 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
}
