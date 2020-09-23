package main

import (
	"fmt"
	"os"

	"github.com/mikestefanello/surveys-microservices/vote-service/serializer"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/config"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/logger"
	"github.com/streadway/amqp"
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

	// Connect to rabbit
	host := fmt.Sprintf("%s:%d", cfg.Rabbit.Hostname, cfg.Rabbit.Port)
	addr := fmt.Sprintf("amqp://%s:%s@%s/", cfg.Rabbit.Username, cfg.Rabbit.Password, host)
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to queue")
		os.Exit(1)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot open channel to queue")
		os.Exit(1)
	}
	defer ch.Close()

	// Initialize queue consumption
	msgs, err := ch.Consume(
		cfg.Rabbit.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to consume queue")
		os.Exit(1)
	}

	// Get a vote serializer
	sz := serializer.NewVoteJSONSerializer()

	// Listen for new messages
	go func() {
		for msg := range msgs {
			v, err := sz.Decode(msg.Body)
			if err != nil {
				log.Error().Err(err).Str("body", string(msg.Body)).Msg("Unable to parse vote from queue message")
				continue
			}
			log.Info().Str("id", v.ID).Msg("Vote received from queue")
		}
	}()

	log.Info().Str("on", host).Msg("Connected to queue. Awaiting messages.")

	// Wait forever
	forever := make(chan bool)
	<-forever
}
