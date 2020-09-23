package main

import (
	"os"

	"github.com/mikestefanello/surveys-microservices/vote-service/serializer"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/config"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/logger"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/queue"
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

	// Get a vote serializer
	sz := serializer.NewVoteJSONSerializer()

	// Load the queue
	mq := queue.NewRabbitVoteQueue(cfg.Rabbit, sz, &log)

	// Listen for new messages
	go func() {
		mq.Consume()
	}()

	// Wait forever
	forever := make(chan bool)
	<-forever
}
