package queue

import (
	"fmt"
	"os"

	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/mikestefanello/surveys-microservices/vote-worker-service/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type RabbitVoteQueue struct {
	cfg        config.RabbitConfig
	serializer vote.Serializer
	log        *zerolog.Logger
}

func NewRabbitVoteQueue(cfg config.RabbitConfig, sz vote.Serializer, l *zerolog.Logger) RabbitVoteQueue {
	return RabbitVoteQueue{
		cfg:        cfg,
		serializer: sz,
		log:        l,
	}
}

func (r *RabbitVoteQueue) Consume() {
	// Connect to rabbit
	host := fmt.Sprintf("%s:%d", r.cfg.Hostname, r.cfg.Port)
	addr := fmt.Sprintf("amqp://%s:%s@%s/", r.cfg.Username, r.cfg.Password, host)
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

	// Initialize queue consumption
	msgs, err := ch.Consume(
		r.cfg.QueueName,
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
		log.Info().Str("id", v.ID).Msg("Vote received from queue")
	}
}
