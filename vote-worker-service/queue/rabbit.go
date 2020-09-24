package queue

import (
	"fmt"
	"os"

	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/config"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type rabbitVoteQueue struct {
	config     config.RabbitConfig
	serializer vote.Serializer
	log        *zerolog.Logger
}

// NewRabbitVoteQueue creates a new rabbit vote queue
func NewRabbitVoteQueue(cfg config.RabbitConfig, sz vote.Serializer, l *zerolog.Logger) VoteQueue {
	return &rabbitVoteQueue{
		config:     cfg,
		serializer: sz,
		log:        l,
	}
}

func (r *rabbitVoteQueue) Consume(vc chan<- *vote.Vote) {
	// Connect to rabbit
	host := fmt.Sprintf("%s:%d", r.config.Hostname, r.config.Port)
	addr := fmt.Sprintf("amqp://%s:%s@%s/", r.config.User, r.config.Password, host)
	conn, err := amqp.Dial(addr)
	if err != nil {
		r.log.Fatal().Err(err).Msg("Cannot connect to queue")
		os.Exit(1)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		r.log.Fatal().Err(err).Msg("Cannot open channel to queue")
		os.Exit(1)
	}
	defer ch.Close()

	// Specify the queue
	_, err = ch.QueueDeclare(
		r.config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.log.Fatal().Err(err).Msg("Cannot declare queue")
		os.Exit(1)
	}

	// Initialize queue consumption
	msgs, err := ch.Consume(
		r.config.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.log.Fatal().Err(err).Msg("Unable to consume queue")
		os.Exit(1)
	}

	r.log.Info().Str("on", host).Msg("Connected to queue. Awaiting messages.")

	for msg := range msgs {
		v, err := r.serializer.Decode(msg.Body)
		if err != nil {
			r.log.Error().Err(err).Str("body", string(msg.Body)).Msg("Unable to parse vote from queue message")
			continue
		}
		r.log.Info().Str("id", v.ID).Msg("Vote received from queue")
		vc <- v
	}
}
